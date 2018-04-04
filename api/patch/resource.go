package patch

import (
	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/fabbricadigitale/scimd/api/query"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/storage"
)

// Resource is ...
func Resource(s storage.Storer, resType *core.ResourceType, id, op, path string, value interface{}) (ret core.ResourceTyper, err error) {

	if resType == nil {
		err = &api.InternalServerError{
			Detail: "ResourceType is nil",
		}
		return ret, err
	}

	p := attr.Parse(path)
	if p.URI == "" {
		p.URI = resType.GetSchema().GetIdentifier()
	}

	ctx := p.Context(resType)
	if ctx == nil {
		err = &api.InvalidPathError{
			Path:   path,
			Detail: "Invalid path error",
		}
		return ret, err
	}

	err = s.Patch(resType, id, "", op, *p, value)
	if err != nil {
		return ret, err
	}

	ret, err = query.Resource(s.(storage.Storer), resType, id, &api.Attributes{})
	return

}
