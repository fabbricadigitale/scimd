package server

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/fabbricadigitale/scimd/config"
	"github.com/fabbricadigitale/scimd/mold"
	"github.com/fabbricadigitale/scimd/validation"

	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/api/messages"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// StaticResourceService describers ...
type StaticResourceService struct {
	resources map[string]interface{}
	endpoint  string
	Service
	Lister
	Getter
}

// NewStaticResourceService creates a new `StaticResourceService` given a path and a list of `core.ResourceTyper`
func NewStaticResourceService(endpoint string, resources interface{}) *StaticResourceService {
	m := make(map[string]interface{})
	rv := reflect.ValueOf(resources)

	switch rv.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < rv.Len(); i++ {
			res := rv.Index(i).Interface()
			m[res.(core.Identifiable).GetIdentifier()] = res
		}
	default:
		panic("not available...")
	}

	return &StaticResourceService{
		endpoint:  endpoint,
		resources: m,
	}
}

// Path returns the endpoint of the `StaticResourceService`
func (rs *StaticResourceService) Path() string {
	return rs.endpoint
}

// List ...
func (rs *StaticResourceService) List(c *gin.Context) {
	// (todo) > return list responses

	params := api.NewSearch()
	// Using the form binding engine (query)
	if err := c.ShouldBindWith(params, binding.Form); err != nil {
		err := messages.NewError(err)
		c.JSON(err.Status, err)
		return
	}

	// (Note) => Can i use MethodNotImplemented middleware instead of the next lines?
	if params.Filter != "" {
		err := messages.NewError(core.ScimError{
			Msg: "Filtering static resource is not supported",
		})
		c.JSON(err.Status, err)
		return
	}

	if e := validation.Validator.Struct(params); e != nil {
		err := messages.NewError(e)
		c.JSON(err.Status, err)
		return
	}

	if e := mold.Transformer.Struct(nil, params); e != nil {
		err := messages.NewError(e)
		c.JSON(err.Status, err)
		return
	}

	var q = &staticQuery{}
	q.Load(rs.resources)

	// Make list
	list := messages.NewListResponse()

	// Count
	list.TotalResults = q.Count()

	if params.Count > config.Values.PageSize {
		params.Count = config.Values.PageSize
	}

	if params.Count == 0 || params.Count > list.TotalResults {
		params.Count = list.TotalResults
	}

	// Pagination
	q.Skip(params.StartIndex - 1).Limit(params.Count - (params.StartIndex - 1))
	list.StartIndex = params.StartIndex
	list.ItemsPerPage = q.Count()

	for _, v := range q.result {
		list.Resources = append(list.Resources, v)
	}

	c.JSON(http.StatusOK, list)
}

// Get ...
func (rs *StaticResourceService) Get(c *gin.Context) {
	id := c.Param("id")
	message := rs.resources[id]
	c.JSON(http.StatusOK, message)
}

type staticQuery struct {
	result []interface{}
}

func (sq *staticQuery) Load(m map[string]interface{}) {
	var s []interface{}
	for _, v := range m {
		s = append(s, v)
	}
	sq.result = s
}

func (sq *staticQuery) Skip(index int) *staticQuery {
	if index < 0 {
		index = 0
	}

	fmt.Printf("index %d", index)

	sq.result = sq.result[index:]
	return sq
}

func (sq *staticQuery) Limit(index int) *staticQuery {
	if index < 0 {
		index = len(sq.result)
	}
	if index > len(sq.result) {
		index = len(sq.result)
	}
	sq.result = sq.result[:index]
	return sq
}

func (sq *staticQuery) Count() int {
	return len(sq.result)
}
