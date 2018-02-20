package validation

import (
	"testing"
	"path/filepath"

	"github.com/stretchr/testify/require"
)

func TestPathExists(t *testing.T){
	var err error
	var absPath string

	okPath := "../../scimd/internal/testdata/service_provider_config.json"
	absPath, _ = filepath.Abs(okPath)
	err = Validator.Var(absPath, "pathexists")
	require.NoError(t, err)

	okPathDir := "../../scimd/internal/testdata"
	absPath, _ = filepath.Abs(okPathDir)
	err = Validator.Var(absPath, "pathexists")
	require.NoError(t, err)

	wrongPath := "../../scimd/internal/testdata/service_provider.json"
	absPath, _ = filepath.Abs(wrongPath)
	err = Validator.Var(absPath, "pathexists")
	require.Error(t, err)

	wrongPathDir := "../../scimd/internal/nonexistingdir"
	absPath, _ = filepath.Abs(wrongPathDir)
	err = Validator.Var(absPath, "pathexists")
	require.Error(t, err)

	emptyPath := ""
	err = Validator.Var(emptyPath, "pathexists")
	require.Error(t, err)

	badType := 123123
	require.PanicsWithValue(t, "Bad field type int", func(){
		err = Validator.Var(badType, "pathexists")
	})
}