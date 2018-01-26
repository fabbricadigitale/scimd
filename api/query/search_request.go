package query

import (
	"github.com/fabbricadigitale/scimd/api/messages"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/storage"
)

func SearchRequest(s storage.Storer, resTypes []*core.ResourceType, search *messages.SearchRequest) (list *messages.ListResponse, err error) {
	// (todo) Validate rew
	return Resources(s, resTypes, &search.Search)
}
