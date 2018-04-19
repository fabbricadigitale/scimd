package patch

import (
	"fmt"

	"github.com/fabbricadigitale/scimd/api/filter"

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

	f, err := getFilter(resType, path)

	err = s.Patch(resType, id, "", op, f, value)
	if err != nil {
		return ret, err
	}

	ret, err = query.Resource(s.(storage.Storer), resType, id, &api.Attributes{})
	return

}

func getFilter(resType *core.ResourceType, path string) (interface{}, error) {
	p := attr.Parse(path)
	if p.URI == "" {
		p.URI = resType.GetSchema().GetIdentifier()
	}

	ctx := p.Context(resType)
	if ctx != nil {
		return p, nil
	}

	f, err := filter.CompileString(path)

	if err != nil {
		err = &api.InvalidPathError{
			Path:   path,
			Detail: "Invalid path error",
		}
		return nil, err
	}
	fn := f.Normalize(resType)
	fmt.Printf("fn %+v\n", fn)
	return fn, nil

}
