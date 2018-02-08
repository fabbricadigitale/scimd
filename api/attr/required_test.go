package attr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateRequired(t *testing.T) {

	err := ValidateRequired(&userRes)
	require.Nil(t, err)

	p := Path{
		Name: "userName",
	}

	p.Context(userRes.ResourceType()).Delete(&userRes)

	err = ValidateRequired(&userRes)
	require.Error(t, err)

}
