package attr

import (
	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/api/messages"
	"github.com/fabbricadigitale/scimd/schemas"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/fabbricadigitale/scimd/schemas/resource"
)

// CheckRequired returns an error if at least one required attribute of res was not asserted by the client
func CheckRequired(res *resource.Resource) (err error) {

	// With the only execption of "schemas", commons required attrs will be not checked because
	// all of them have "readOnly" mutability which attrs cannot be asserted by the client
	// and must be ignored.

	if len(res.Schemas) == 0 {
		return messages.NewError(&api.MissingRequiredPropertyError{
			Path: "schemas",
		})
	}

	attrs, err := Paths(res.ResourceType(), func(attribute *core.Attribute) bool {
		return (attribute.Required == true && attribute.Mutability != schemas.MutabilityReadOnly)
	})

	if err != nil {
		return
	}

	for _, a := range attrs {

		ctx := a.Context(res.ResourceType())

		if ctx.Schema == nil {
			continue
		}

		v := ctx.Get(res)

		if datatype.IsNull(v) {
			err = &api.MissingRequiredPropertyError{
				Path: a.String(),
			}
		}

	}
	return
}
