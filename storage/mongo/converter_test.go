package mongo

import (
	"testing"

	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/stretchr/testify/require"

	"github.com/globalsign/mgo/bson"
)

type tCase struct {
	path  string
	value []interface{}
	op    string
	query bson.M
}

var tCases = []tCase{
	{
		path: "emails",
		value: []interface{}{
			datatype.Complex{
				"value": "gigi@gmail.com",
				"type":  "private",
			},
		},
		op:    "add",
		query: bson.M{"$push": bson.M{"urn:ietf:params:scim:schemas:core:2°0:User.emails": []interface{}{bson.M{"type": "private", "value": "gigi@gmail.com"}}}},
	},
	{
		path: "emails",
		value: []interface{}{
			datatype.Complex{
				"value": "gigi@gmail.com",
				"type":  "private",
			},
		},
		op:    "replace",
		query: bson.M{"$set": bson.M{"urn:ietf:params:scim:schemas:core:2°0:User.emails": []interface{}{bson.M{"type": "private", "value": "gigi@gmail.com"}}}},
	},
	{
		path: "emails",
		value: []interface{}{
			datatype.Complex{
				"value": "gigi@gmail.com",
				"type":  "private",
			},
		},
		op:    "remove",
		query: bson.M{"$pull": bson.M{"urn:ietf:params:scim:schemas:core:2°0:User.emails": []interface{}{bson.M{"value": "gigi@gmail.com", "type": "private"}}}},
	},
}

func TestConvertChangeValue(t *testing.T) {

	//repo := core.GetResourceTypeRepository()

	for _, tc := range tCases {
		query, err := convertChangeValue(resTypeRepo.Pull("User"), tc.op, attr.Path{
			URI:  "urn:ietf:params:scim:schemas:core:2.0:User",
			Name: "emails",
		}, tc.value)
		require.NoError(t, err)
		require.Equal(t, tc.query, query)
	}

}
