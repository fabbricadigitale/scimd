package schemas

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenericRepository(t *testing.T) {
	repo1 := GetGenericRepository()
	repo2 := GetGenericRepository()

	// Is it a singleton?
	require.IsType(t, (*repositoryGeneric)(nil), repo1)
	require.IsType(t, (*repositoryGeneric)(nil), repo2)
	require.Implements(t, (*GenericRepository)(nil), repo1)
	require.Implements(t, (*GenericRepository)(nil), repo2)
	require.Exactly(t, repo1, repo2)
}

func TestSchemaRepository(t *testing.T) {
	// (todo)
}

func TestResourceTypeRepository(t *testing.T) {
	// (todo)
}
