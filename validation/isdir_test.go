package validation

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsDir(t *testing.T){
	var err error
	var absPath string

	existingDirPath := "../../scimd/internal/testdata"
	absPath, _ = filepath.Abs(existingDirPath)
	err = Validator.Var(absPath, "isdir")
	require.NoError(t, err)

	nonExistingDirPath := "../../scimd/internal/nonexistingdir"
	absPath, _ = filepath.Abs(nonExistingDirPath)
	err = Validator.Var(absPath, "isdir")
	require.Error(t, err)

	filePath := "../../scimd/internal/testdata/service_provider_config.json"
	absPath, _ = filepath.Abs(filePath)
	err = Validator.Var(absPath, "isdir")
	require.Error(t, err)

	emptyPath := ""
	err = Validator.Var(emptyPath, "isdir")
	require.Error(t, err)

	badType := 123123
	require.PanicsWithValue(t, "Bad field type int", func(){
		err = Validator.Var(badType, "isdir")
	})
}