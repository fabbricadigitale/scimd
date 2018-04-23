package mongo

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/fabbricadigitale/scimd/api/attr"

	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/api/filter"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/globalsign/mgo/bson"
)

func makeComplexQuery(resType *core.ResourceType, id, version string, f interface{}) (bson.M, error) {

	if f == nil {
		return makeQuery(resType.GetIdentifier(), id, version), nil
	}

	switch f.(type) {
	case string:
		return makeQuery(resType.GetIdentifier(), id, version), nil
	case filter.Filter:
		return convertToMongoQuery(resType, f.(filter.Filter), "id", id)
	default:
		return makeQuery(resType.GetIdentifier(), id, version), nil
	}
}

func getPath(f interface{}) attr.Path {

	switch f.(type) {
	case *attr.Path:
		return *(f.(*attr.Path))
	case *filter.ValuePath:
		node := f.(*filter.ValuePath)
		return node.Path
	default:
		return attr.Path{}
	}
}

func (c *convert) extractPath(f interface{}) attr.Path {

	switch f.(type) {
	case *attr.Path:
		return *(f.(*attr.Path))
	case *filter.Group:
		node := f.(*filter.Group)
		return c.extractPath(node.Filter)
	case *filter.ValuePath:
		node := f.(*filter.ValuePath)
		return node.Path
	case *filter.AttrExpr:
		node := f.(*filter.AttrExpr)
		return node.Path
	default:
		return attr.Path{}
	}

}

func convertChangeValue(resType *core.ResourceType, op string, p attr.Path, value interface{}) (m bson.M, err error) {

	if resType == nil {
		err = &api.InternalServerError{
			Detail: "ResourceType is nil",
		}
		return nil, err
	}

	if p.Undefined() {
		err = &api.InternalServerError{
			Detail: "Path is undefined",
		}
		return nil, err
	}

	if value == nil && (op == "add" || op == "replace") {
		err = &api.InternalServerError{
			Detail: "Value is nil",
		}
		return nil, err
	}

	ctx := p.Context(resType)

	if ctx.Attribute.MultiValued == false {
		m = getBSONSingleValued(op, p, value)
	} else {
		m = getBSONMultiValued(op, p, value)
	}

	return m, err
}

func getBSONSingleValued(op string, p attr.Path, value interface{}) bson.M {

	path := escapeAttribute(p.String())

	var operator string

	if op == "add" || op == "replace" {
		operator = "$set"
	} else {
		// remove
		operator = "$unset"
	}

	m := bson.M{}
	ret := bson.M{}

	switch value.(type) {

	case map[string]interface{}:
		values := value.(map[string]interface{})

		o := bson.M{}

		for key, val := range values {
			o[escapeAttribute(key)] = val
		}

		m = bson.M{
			path: o,
		}

		break
	default:
		m = bson.M{
			path: value,
		}

		break
	}

	ret[operator] = m

	return ret
}

func getBSONMultiValued(op string, p attr.Path, value interface{}) bson.M {

	path := escapeAttribute(p.String())

	var operator string
	modifier := ""

	if op == "add" {
		operator = "$push"
		modifier = "$each"

	} else if op == "remove" {
		if value != nil {
			operator = "$pull"
		} else {
			operator = "$unset"
		}

	} else if op == "replace" {
		operator = "$set"
		path = escapeAttribute(fmt.Sprintf("%s:%s", p.URI, p.Name))

		v := reflect.ValueOf(value)
		if !(v.Kind() == reflect.Array || v.Kind() == reflect.Slice) {
			path = fmt.Sprintf("%s.$", path)
		}
	}

	m := bson.M{}
	ret := bson.M{}

	if value != nil {
		switch value.(type) {

		case map[string]interface{}:
			o := escapeValue(value.(map[string]interface{}))
			m = bson.M{path: o}
			break

		case []interface{}:
			values := value.([]interface{})

			o := make([]bson.M, 0)
			for _, v := range values {
				i := escapeValue(v.(map[string]interface{}))
				o = append(o, i)
			}
			if modifier != "" {
				e := bson.M{}
				e[modifier] = o
				m = bson.M{path: e}
			} else {
				m = bson.M{path: o}
			}

			break

		default:
			m = bson.M{path: value}
			break
		}
	} else {
		m = bson.M{path: ""}
	}

	ret[operator] = m

	return ret
}

func escapeValue(value map[string]interface{}) bson.M {
	o := bson.M{}

	for key, val := range value {
		o[escapeAttribute(key)] = val
	}

	return o
}

func convertToMongoQuery(resType *core.ResourceType, ft filter.Filter, key string, value interface{}) (m bson.M, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case error:
				err = r.(error)
			default:
				err = &api.InternalServerError{
					Detail: r.(string),
				}
			}
		}
	}()

	var normFilterByResType filter.Filter
	if ft != nil {
		normFilterByResType = ft.Normalize(resType)
	}

	var conv *convert
	m = conv.do(resType, normFilterByResType)
	m[key] = value
	return m, err
}

type convert struct{}

func (c *convert) do(resType *core.ResourceType, f interface{}) bson.M {
	var (
		left, right bson.M
	)

	switch f.(type) {
	case *filter.Group:
		node := f.(*filter.Group)
		return c.do(resType, node.Filter)

	case *filter.And:
		node := f.(*filter.And)
		if node.Left != nil {
			left = c.do(resType, node.Left)
		}
		if node.Right != nil {
			right = c.do(resType, node.Right)
		}
		return bson.M{
			"$and": []interface{}{left, right},
		}
	case *filter.Or:
		node := f.(*filter.Or)
		if node.Left != nil {
			left = c.do(resType, node.Left)
		}
		if node.Right != nil {
			right = c.do(resType, node.Right)
		}
		return bson.M{
			"$or": []interface{}{left, right},
		}
	case *filter.Not:
		node := f.(*filter.Not)
		left = c.do(resType, node.Filter)
		return bson.M{
			"$nor": []interface{}{left},
		}
	case *filter.AttrExpr:
		node := f.(*filter.AttrExpr)
		return c.relationalOperators(resType, f, node)
	default:
		return bson.M{}
	}
}

// Represent a mongo key that's always not present
const notExistingKey = "_"

func (c *convert) relationalOperators(resType *core.ResourceType, f interface{}, node *filter.AttrExpr) bson.M {
	// If any schema attribure was not found node.Value is nil.
	// For filtered attributes that are not part of a particular resource
	// type, the service provider SHALL treat the attribute as if there is
	// no attribute value, as per https://tools.ietf.org/html/rfc7644#section-3.4.2.1
	if node.Path.Undefined() {
		return bson.M{
			notExistingKey: bson.M{
				mapOperator[node.Op]: node.Value,
			},
		}
	}

	// The 'co', 'sw' and 'ew' operators can only be used if the attribute type is string
	if node.Op == filter.OpContains || node.Op == filter.OpStartsWith || node.Op == filter.OpEndsWith {
		return stringOperators(resType, f, node)
	} else if node.Op == filter.OpPresent {
		return prOperator(resType, f, node)
	} else {
		return comparisonOperators(resType, f, node)
	}
}

func newInvalidFilterError(detail, filter string) *api.InvalidFilterError {
	var e *api.InvalidFilterError
	e = &api.InvalidFilterError{
		Filter: filter,
		Detail: detail,
	}
	return e
}

func stringOperators(resType *core.ResourceType, f interface{}, node *filter.AttrExpr) bson.M {
	key := pathToKey(node.Path)
	value := node.Value.(string)

	switch node.Op {
	case filter.OpContains:
		return regexQueryPart(key, value, "i", "", "")
	case filter.OpStartsWith:
		return regexQueryPart(key, value, "i", "^", "")
	case filter.OpEndsWith:
		return regexQueryPart(key, value, "i", "", "$")
	default:
		return nil
	}
}

func regexQueryPart(key, value, option, prePattern, postPattern string) bson.M {
	return bson.M{
		key: bson.M{
			"$regex": bson.RegEx{
				Pattern: prePattern + regexp.QuoteMeta(value) + postPattern,
				Options: option,
			},
		},
	}
}

func comparisonOperators(resType *core.ResourceType, f interface{}, node *filter.AttrExpr) bson.M {
	key := pathToKey(node.Path)
	return bson.M{
		key: bson.M{
			mapOperator[node.Op]: node.Value,
		},
	}

}

func prOperator(resType *core.ResourceType, f interface{}, node *filter.AttrExpr) bson.M {

	attrDef := node.Path.Context(resType).Attribute
	key := pathToKey(node.Path)

	existsCriteria := bson.M{key: bson.M{"$exists": true}}
	nullCriteria := bson.M{key: bson.M{"$ne": nil}}
	emptyStringCriteria := bson.M{key: bson.M{"$ne": ""}}
	emptyArrayCriteria := bson.M{key: bson.M{"$not": bson.M{"$size": 0}}}
	emptyObjectCriteria := bson.M{key: bson.M{"$ne": bson.M{}}}

	criterion := make([]interface{}, 0)
	criterion = append(criterion, existsCriteria, nullCriteria)
	if attrDef.MultiValued {
		criterion = append(criterion, emptyArrayCriteria)
	} else {
		switch attrDef.Type {
		case datatype.StringType:
			criterion = append(criterion, emptyStringCriteria)
		case datatype.ComplexType:
			criterion = append(criterion, emptyObjectCriteria)
		}
	}
	return bson.M{
		key: bson.M{"$and": criterion},
	}
}
