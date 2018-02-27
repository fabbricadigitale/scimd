package attributes

import (
	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/fabbricadigitale/scimd/schemas"
	"github.com/fabbricadigitale/scimd/schemas/core"
)

// GetUniqueAttributes is ...
// (todo) > add support to multi key indexes
func GetUniqueAttributes() ([][]string, error) {

	list := core.GetResourceTypeRepository().List()

	var uniqueAttrs [][]string
	uniqueAttrs = make([][]string, 0)

	for _, resType := range list {

		ro, err := attr.Paths(&resType, func(attribute *core.Attribute) bool {
			return attribute.Uniqueness == schemas.UniquenessServer || attribute.Uniqueness == schemas.UniquenessGlobal
		})
		if err != nil {
			return nil, err
		}

		for _, p := range ro {
			uniqueAttrs = append(uniqueAttrs, []string{p.String()})
		}

	}

	return uniqueAttrs, nil
}
