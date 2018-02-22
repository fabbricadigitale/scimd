package mongo

import (
	"regexp"

	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/api/filter"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/globalsign/mgo/bson"
)

func convertToMongoQuery(resType *core.ResourceType, ft filter.Filter) (m bson.M, err error) {
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
	m["meta.resourceType"] = resType.GetIdentifier()
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
