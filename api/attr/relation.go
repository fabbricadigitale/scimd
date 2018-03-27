package attr

import (
	"github.com/fabbricadigitale/scimd/schemas"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
)

// Relation is ...
type Relation struct {
	RWAttribute    Path
	ROAttribute    Path
	ROResourceType core.ResourceType
}

// GetRelationships is ...
func GetRelationships(key, resourceTypeID string) ([]Relation, error) {

	ret := make([]Relation, 0)

	repo := core.GetSchemaRepository()
	resTypeRepo := core.GetResourceTypeRepository()
	schema := repo.Pull(key)

	if schema == nil {
		return nil, core.ScimError{
			Msg: "Error Schema is nil",
		}
	}

	for _, attribute := range (*schema).Attributes {

		if (*attribute).Type == datatype.ComplexType && (*attribute).Mutability == schemas.MutabilityReadWrite {

			for _, subAttr := range (*attribute).SubAttributes {

				if (*subAttr).Type == datatype.ReferenceType {

					if len(subAttr.ReferenceTypes) > 0 {

						referenceType := subAttr.ReferenceTypes[0]

						if referenceType == "external" {
							continue
						}

						roResType := resTypeRepo.Pull(referenceType)

						roSchema := repo.Pull(roResType.Schema)

						for _, roAttribute := range (*roSchema).Attributes {

							if roAttribute.Type == datatype.ComplexType && roAttribute.Mutability == schemas.MutabilityReadOnly {

								for _, roSubAttr := range (*roAttribute).SubAttributes {

									if (*roSubAttr).Type == datatype.ReferenceType {

										if referenceType == "external" {
											continue
										}

										if len(roSubAttr.ReferenceTypes) > 0 && roSubAttr.ReferenceTypes[0] == resourceTypeID {

											r := Relation{
												RWAttribute: Path{
													URI:  key,
													Name: (*attribute).Name,
												},
												ROAttribute: Path{
													URI:  roSchema.ID,
													Name: (*roAttribute).Name,
												},
												ROResourceType: *roResType,
											}
											ret = append(ret, r)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return ret, nil

}
