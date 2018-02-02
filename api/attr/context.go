package attr

import (
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/fabbricadigitale/scimd/schemas/resource"
)

// A Context represents a set of definitions related to a Path
type Context struct {
	Schema       *core.Schema
	Attribute    *core.Attribute
	SubAttribute *core.Attribute
}

// Path returns a Path built from a given Context
func (ctx *Context) Path() *Path {
	p := Path{}

	if ctx.Schema != nil {
		p.URI = ctx.Schema.ID
	}

	if ctx.Attribute != nil {
		p.Name = ctx.Attribute.Name
	}

	if ctx.SubAttribute != nil {
		p.Sub = ctx.SubAttribute.Name
	}

	return &p
}

// Context fetches from rt a suitable Context for p, if any.
func (p Path) Context(rt *core.ResourceType) (ctx *Context) {

	// (todo) implement caching

	if rt == nil || p.Undefined() {
		return
	}

	ctx = &Context{}

	// Try common attributes
	if p.URI == "" {
		ctx.Attribute = core.Commons().WithName(p.Name)
		if ctx.Attribute != nil {
			if p.Sub != "" {
				ctx.SubAttribute = ctx.Attribute.SubAttributes.WithName(p.Sub)
				if ctx.SubAttribute == nil {
					// Unmached path
					return nil
				}
			}
			return
		}
	}

	// Try schema attributes
	ctx.Schema = p.matchSchema(rt)
	if ctx.Schema == nil {
		// Unmached path
		return nil
	}

	ctx.Attribute = ctx.Schema.Attributes.WithName(p.Name)
	if ctx.Attribute != nil {
		if p.Sub != "" {
			ctx.SubAttribute = ctx.Attribute.SubAttributes.WithName(p.Sub)
			if ctx.SubAttribute == nil {
				// Unmached path
				return nil
			}
		}
	}

	return
}

// Contexts returns a slice of Context given a resource type rt.
//
// It flattens the attributes of rt's schemas returning their Context representations.
// When a fx is provided it returns only the attribute statisfying fx(attribute).
func Contexts(rt *core.ResourceType, fx func(attribute *core.Attribute) bool) []Context {
	// Tautology
	if fx == nil {
		fx = func(attribute *core.Attribute) bool {
			return true
		}
	}

	// Accumulation iterating over all contexts
	acc := []Context{}

	commonCtx := Context{} // Common attributes have no schema
	for _, c1 := range core.Commons() {
		commonCtx.Attribute = c1
		commonCtx.SubAttribute = nil
		if fx(c1) {
			acc = append(acc, commonCtx)
		}
		for _, c2 := range c1.SubAttributes {
			commonCtx.SubAttribute = c2
			if fx(c2) {
				acc = append(acc, commonCtx)
			}
		}
	}

	for _, sc := range rt.GetSchemas() {
		if sc != nil {
			ctx := Context{
				Schema: sc,
			}
			for _, a1 := range sc.Attributes {
				ctx.Attribute = a1
				ctx.SubAttribute = nil
				if fx(a1) {
					acc = append(acc, ctx)
				}
				for _, a2 := range a1.SubAttributes {
					ctx.SubAttribute = a2
					if fx(a2) {
						acc = append(acc, ctx)
					}
				}
			}
		}
	}

	return acc
}

func (ctx *Context) getValuerValues(valuer resource.Valuer) *datatype.Complex {
	if ctx.Schema == nil {
		return nil
	}
	values := valuer.Values(ctx.Schema.GetIdentifier())
	if values == nil {
		values = &datatype.Complex{}
		valuer.SetValues(ctx.Schema.GetIdentifier(), values)
	}
	return values
}

// Set value to valuer at destination path of this context
//
// Value should be datatype.DataTyper or []datatype.DataTyper or nil.
// Furthermore attribute's characteristics are not enforced nor checked
// according to DataTypes rules when used within maps.
//
// However, when context is pointing to a sub-attribute and parent attribute is missing,
// parent's value will be created with respect of attribute definition in order to accomodate
// the nested sub-attribute's value.
//
// Value will be not set if it's not possibile to accomodate the sub-attribute's value in any way
// (eg. trying to set a sub-attribute's value in a map that's not respecting the parent attribute's definition).
//
// Usage example:
// Parse("emails.type").Context(res.ResourceType()).Set(datatype.String("home"), res)
// for each complex value within "emails" (it's a multi value), the value of "type" sub-attribute will be set to "home".
func (ctx *Context) Set(value interface{}, valuer resource.Valuer) {
	values := ctx.getValuerValues(valuer)
	if values == nil || ctx.Attribute == nil {
		return
	}

	if ctx.SubAttribute == nil {
		(*values)[ctx.Attribute.Name] = value
		return
	}

	// Make parent if missing
	if (*values)[ctx.Attribute.Name] == nil && ctx.Attribute.Type == datatype.ComplexType {
		if ctx.Attribute.MultiValued {
			(*values)[ctx.Attribute.Name] = []datatype.DataTyper{&datatype.Complex{}}
		} else {
			(*values)[ctx.Attribute.Name] = &datatype.Complex{}
		}
	}

	if ctx.Attribute.MultiValued {
		if parent, ok := (*values)[ctx.Attribute.Name].([]datatype.DataTyper); ok {
			for _, vv := range parent {
				if elem, ok := vv.(datatype.Complex); ok {
					elem[ctx.SubAttribute.Name] = value
				}
			}
		}
	} else {
		if parent, ok := (*values)[ctx.Attribute.Name].(datatype.Complex); ok {
			parent[ctx.SubAttribute.Name] = value
		}
	}
}

// Get a value from valuer at destination path of this context and return it
func (ctx *Context) Get(valuer resource.Valuer) interface{} {
	values := ctx.getValuerValues(valuer)
	if values == nil || ctx.Attribute == nil {
		return nil
	}

	if ctx.SubAttribute == nil {
		return (*values)[ctx.Attribute.Name]
	}

	if ctx.Attribute.MultiValued {
		if parent, ok := (*values)[ctx.Attribute.Name].([]datatype.DataTyper); ok {
			ret := make([]interface{}, len(parent))
			for i, vv := range parent {
				if elem, ok := vv.(datatype.Complex); ok {
					ret[i] = elem[ctx.SubAttribute.Name]
				}
			}
			return ret
		}
	} else {
		if parent, ok := (*values)[ctx.Attribute.Name].(datatype.Complex); ok {
			return parent[ctx.SubAttribute.Name]
		}
	}

	return nil
}

// Delete a value from valuer at destination path of this context
func (ctx *Context) Delete(valuer resource.Valuer) {
	values := ctx.getValuerValues(valuer)
	if values == nil || ctx.Attribute == nil {
		return
	}

	if ctx.SubAttribute == nil {
		delete((*values), ctx.Attribute.Name)
		return
	}

	if ctx.Attribute.MultiValued {
		if parent, ok := (*values)[ctx.Attribute.Name].([]datatype.DataTyper); ok {
			for _, vv := range parent {
				if elem, ok := vv.(datatype.Complex); ok {
					delete(elem, ctx.SubAttribute.Name)
				}
			}
		}
	} else {
		if parent, ok := (*values)[ctx.Attribute.Name].(datatype.Complex); ok {
			delete(parent, ctx.SubAttribute.Name)
		}
	}
}
