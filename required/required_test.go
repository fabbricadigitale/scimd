package required

import (
	"testing"

	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/stretchr/testify/require"
)

func TestValidateRequired(t *testing.T) {

	err := ValidateRequired(&userRes)
	require.Nil(t, err)

	p := attr.Path{
		Name: "userName",
	}

	p.Context(userRes.ResourceType()).Delete(&userRes)

	err = ValidateRequired(&userRes)
	require.Error(t, err)

}
