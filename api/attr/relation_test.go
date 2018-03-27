package attr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	key            string
	resourceTypeID string
	expected       []Relation
}

func TestGetRelationships(t *testing.T) {

	testCases := []testCase{
		{
			key:            "urn:ietf:params:scim:schemas:core:2.0:Group",
			resourceTypeID: "Group",
			expected: []Relation{
				Relation{
					RWAttribute:    Path{URI: "urn:ietf:params:scim:schemas:core:2.0:Group", Name: "members", Sub: ""},
					ROAttribute:    Path{URI: "urn:ietf:params:scim:schemas:core:2.0:User", Name: "groups", Sub: ""},
					ROResourceType: *resTypeRepo.Pull("User"),
				},
			},
		},
	}

	for _, item := range testCases {
		ret, _ := GetRelationships(item.key, item.resourceTypeID)
		require.Equal(t, item.expected, ret)

	}
}
