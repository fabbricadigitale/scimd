package required

import (
	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/fabbricadigitale/scimd/api/messages"
	"github.com/fabbricadigitale/scimd/schemas"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/fabbricadigitale/scimd/validation"
)

// ValidateRequired is a function to check that required field is not null
func ValidateRequired(res *resource.Resource) (err error) {

	attrs := attr.Paths(res.ResourceType(), func(attribute *core.Attribute) bool {
		return (attribute.Required == true && attribute.Mutability != schemas.MutabilityReadOnly)
	})

	// ID is a required but readOnly property.
	// Using StructExcept we can exclude ID attribute validation.
	err = validation.Validator.StructExcept(res.CommonAttributes, "ID")
	if err != nil {
		err = messages.NewError(&api.MissingRequiredPropertyError{
			Detail: err.Error(),
		})
		return
	}

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
