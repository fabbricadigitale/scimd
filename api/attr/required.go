package attr

import (
	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/api/messages"
	"github.com/fabbricadigitale/scimd/schemas"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/fabbricadigitale/scimd/schemas/resource"
)

// ValidateRequired is a function to check that required field is not null
func ValidateRequired(res *resource.Resource) (err error) {

	attrs := Paths(res.ResourceType(), func(attribute *core.Attribute) bool {
		return (attribute.Required == true && attribute.Mutability != schemas.MutabilityReadOnly)
	})

	// TODO: Checks required in CommonAttribute

	for _, attr := range attrs {

		ctx := attr.Context(res.ResourceType())

		if ctx.Schema == nil {
			continue
		}

		x := ctx.Get(res)

		if datatype.IsNull(x) {
			err = messages.NewError(&api.MissingRequiredPropertyError{
				Path: attr.String(),
			})
		}

	}
	return
}
