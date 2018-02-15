package harness

import (
	"net/http/httptest"
	"testing"

	scim "github.com/fabbricadigitale/scimd/cmd/scimd"
)

func TestSimpleGet(t *testing.T) {
	setup()
	defer teardown()

	eng := scim.GetEngine()
	rec := httptest.NewRecorder()
}
