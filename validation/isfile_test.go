package validation

import (
	"github.com/stretchr/testify/require"
	"path/filepath"
	"testing"
)

func TestIsFile(t *testing.T){
	var err error
	var absPath string

	existingFile := "../../scimd/internal/testdata/service_provider_config.json"
	absPath, _ = filepath.Abs(existingFile)
	err = Validator.Var(absPath, "isfile")
	require.NoError(t, err)

	nonExistingFile := "../../scimd/internal/testdata/service_provider.json"
	absPath, _ = filepath.Abs(nonExistingFile)
	err = Validator.Var(absPath, "isfile")
	require.Error(t, err)

	dirPath := "../../scimd/internal/testdata"
	absPath, _ = filepath.Abs(dirPath)
	err = Validator.Var(absPath, "isfile")
	require.Error(t, err)

	okFileExtension := "../../scimd/internal/testdata/service_provider_config.json"
	err = Validator.Var(okFileExtension, "isfile=json")
	require.NoError(t, err)

	okFileExtension2 := "../../scimd/internal/test/integration/resource_type_test.go"
	err = Validator.Var(okFileExtension2, "isfile=go")
	require.NoError(t, err)

	wrongFileExtension := "../../scimd/internal/testdata/service_provider_config.json"
	err = Validator.Var(wrongFileExtension, "isfile=go")
	require.Error(t, err)

	dirExtension := "../../scimd/internal/testdata"
	err = Validator.Var(dirExtension, "isfile=dir")
	require.Error(t, err)

	empty := ""
	err = Validator.Var(empty, "isfile")
	require.Error(t, err)

	badType := 123123
	require.PanicsWithValue(t, "Bad field type int", func(){
		err = Validator.Var(badType, "isfile")
	})
}