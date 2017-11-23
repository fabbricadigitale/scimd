package filter

import (
	"testing"

	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/stretchr/testify/assert"
)

const (
	filter1 = `userType ne "Employee" and not (emails co "example.com" or emails.value co "example.org")`
)

func TestStringer(t *testing.T) {
	var f1 Filter = And{
		AttrExpr{attr.Path{Name: "userType"}, OpNotEqual, "Employee"},
		Not{
			Or{
				AttrExpr{attr.Path{Name: "emails"}, OpContains, "example.com"},
				AttrExpr{attr.Path{Name: "emails", Sub: "value"}, OpContains, "example.org"},
			},
		},
	}

	assert.Equal(t, filter1, f1.String())
}
