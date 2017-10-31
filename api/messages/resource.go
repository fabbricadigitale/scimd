package messages

import (
	"github.com/fabbricadigitale/scimd/schemas/core"
)

// Resource The data resource structure
type Resource struct {
	Common     core.Common
	Attributes map[string]core.Complex
}
