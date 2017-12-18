package mongo

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"testing"

	"github.com/fabbricadigitale/scimd/api/filter"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/stretchr/testify/require"
	"gopkg.in/mgo.v2/bson"
)

type mongoCase struct {
	filter string // the input filter
	// query1 bson.M // the query with unknown attributes (ie., no resource type and schema repository)
	query2 bson.M // the complete query
}

var mongoTests = []mongoCase{
	{
		`emails[type eq "work" and value co "@example.com"]`,
		bson.M{"meta.resourceType": "User", "$and": []interface{}{bson.M{"_": bson.M{"$eq": "work"}}, bson.M{"_": bson.M{"": "@example.com"}}}},
		bson.M{"$and": []interface{}{bson.M{"urn:ietf:params:scim:schemas:core:2°0:User.emails.type": bson.M{"$eq": "work"}}, bson.M{"urn:ietf:params:scim:schemas:core:2°0:User.emails.value": bson.M{"$regex": bson.RegEx{Pattern: "@example\\.com", Options: "i"}}}}, "meta.resourceType": "User"},
	},
	{
		`emails[type eq "work" and value co "@example.com"] or ims[type eq "xmpp" and value co "@foo.com"]`,
		bson.M{"$or": []interface{}{bson.M{"$and": []interface{}{bson.M{"_": bson.M{"$eq": "work"}}, bson.M{"_": bson.M{"": "@example.com"}}}}, bson.M{"$and": []interface{}{bson.M{"_": bson.M{"$eq": "xmpp"}}, bson.M{"_": bson.M{"": "@foo.com"}}}}}, "meta.resourceType": "User"},
		bson.M{"$or": []interface{}{bson.M{"$and": []interface{}{bson.M{"urn:ietf:params:scim:schemas:core:2°0:User.emails.type": bson.M{"$eq": "work"}}, bson.M{"urn:ietf:params:scim:schemas:core:2°0:User.emails.value": bson.M{"$regex": bson.RegEx{Pattern: "@example\\.com", Options: "i"}}}}}, bson.M{"$and": []interface{}{bson.M{"urn:ietf:params:scim:schemas:core:2°0:User.ims.type": bson.M{"$eq": "xmpp"}}, bson.M{"urn:ietf:params:scim:schemas:core:2°0:User.ims.value": bson.M{"$regex": bson.RegEx{Pattern: "@foo\\.com", Options: "i"}}}}}}, "meta.resourceType": "User"},
	},
	/*
		// (todo) > complete test cases .. good luck
			{
				`userType eq "Employee" and emails[type eq "work" and value co "@example.com"]`,
				``,
				``,
			},
			{
				`title pr`,
				``,
				``,
			},
			{
				`emails[not (type sw null)]`,
				``,
				``,
			},
			{
				`userType eq "Employee" and (emails co "example.com" or emails.value co "example.org")`,
				``,
				``,
			},
			{
				`userType eq "Employee" and emails.type eq "work"`,
				``,
				``,
			},
			{
				`userType eq "Employee" and (emails.type eq "work")`,
				``,
				``,
			},
			{
				`emails.type eq "work"`,
				``,
				``,
			},
			{
				`userName eq "bjensen" and name.familyName sw "J"`,
				``,
				``,
			},
			{
				`not (userName.Child eq "strings")`,
				``,
				``,
			},
			{
				`userName sw "J"`,
				``,
				``,
			},
			{
				`emails co "example.com"`,
				``,
				``,
			},
			{
				`emails.type co "work"`,
				``,
				``,
			},
			{
				`emails.type ne true`,
				``,
				``,
			},
			{
				`userType ne "Employee" and not (emails co "example.com" or emails.value co "example.org")`,
				``,
				``,
			},
			{
				`userName eq "bjensen" and name.familyName sw "J"`,
				``,
				``,
			},
			{
				`userType ne "Employee" and not (emails co "example.com" or emails.value co "example.org")`,
				``,
				``,
			},
			{
				`userName eq "bjensen"`,
				``,
				``,
			},
			{
				`meta.lastModified gt "2011-05-13T04:42:34Z"`,
				``,
				``,
			},
			{
				`meta.lastModified ge "2011-05-13T04:42:34Z"`,
				``,
				``,
			},
			{
				`meta.lastModified lt "2011-05-13T04:42:34Z"`,
				``,
				``,
			},
			{
				`meta.lastModified le "2011-05-13T04:42:34Z"`,
				``,
				``,
			},
			{
				`name.familyName co "O'Malley"`,
				``,
				``,
			},
			{
				`not (userName eq "strings")`,
				``,
				``,
			},
	*/
		// bson.M{"meta.resourceType": "User", "$and": []interface{}{bson.M{"_": bson.M{"$eq": "work"}}, bson.M{"_": bson.M{"": "@example.com"}}}},
		bson.M{"$and": []interface{}{
			bson.M{"$elemMatch": bson.M{"$and": []interface{}{
				bson.M{"emails.type": bson.M{"$eq": "work"}}, bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}}},
			bson.M{"$elemMatch": bson.M{"$and": []interface{}{
				bson.M{"emails.value": bson.M{"$regex": bson.RegEx{Pattern: "@example\\.com", Options: "i"}}},
				bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}}}},
			"meta.resourceType": "User",
		},
	},
	{
		`emails[type eq "work" and value co "@example.com"] or ims[type eq "xmpp" and value co "@foo.com"]`,
		// bson.M{"$or": []interface{}{bson.M{"$and": []interface{}{bson.M{"_": bson.M{"$eq": "work"}}, bson.M{"_": bson.M{"": "@example.com"}}}}, bson.M{"$and": []interface{}{bson.M{"_": bson.M{"$eq": "xmpp"}}, bson.M{"_": bson.M{"": "@foo.com"}}}}}, "meta.resourceType": "User"},
		bson.M{"$or": []interface{}{
			bson.M{"$and": []interface{}{
				bson.M{"$elemMatch": bson.M{"$and": []interface{}{
					bson.M{"emails.type": bson.M{"$eq": "work"}},
					bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}}},
				bson.M{"$elemMatch": bson.M{"$and": []interface{}{
					bson.M{"emails.value": bson.M{"$regex": bson.RegEx{Pattern: "@example\\.com", Options: "i"}}},
					bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}}}}},
			bson.M{"$and": []interface{}{
				bson.M{"$elemMatch": bson.M{"$and": []interface{}{
					bson.M{"ims.type": bson.M{"$eq": "xmpp"}},
					bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}}},
				bson.M{"$elemMatch": bson.M{"$and": []interface{}{
					bson.M{"ims.value": bson.M{"$regex": bson.RegEx{Pattern: "@foo\\.com", Options: "i"}}},
					bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}}}}}},
			"meta.resourceType": "User",
		},
	},
	{
		`userType eq "Employee" and emails[type eq "work" and value co "@example.com"]`,

		bson.M{"$and": []interface{}{
			bson.M{"$and": []interface{}{
				bson.M{"userType": bson.M{"$eq": "Employee"}},
				bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}},
			bson.M{"$and": []interface{}{
				bson.M{"$elemMatch": bson.M{"$and": []interface{}{
					bson.M{"emails.type": bson.M{"$eq": "work"}}, bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}}},
				bson.M{"$elemMatch": bson.M{"$and": []interface{}{
					bson.M{"emails.value": bson.M{"$regex": bson.RegEx{Pattern: "@example\\.com", Options: "i"}}},
					bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}}}}}},
			"meta.resourceType": "User",
		},
	},
	{
		`title pr`,
		bson.M{"$and": []interface{}{
			bson.M{"$and": []interface{}{
				bson.M{"title": bson.M{"$exists": true}},
				bson.M{"title": bson.M{"$ne": interface{}(nil)}},
				bson.M{"title": bson.M{"$ne": ""}}}},
			bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}},
			"meta.resourceType": "User",
		},
	},
	/* 	{
		`emails[not (type sw null)]`, // TO CHECK - error: interface conversion: interface {} is nil, not string
	}, */
	{
		`userType eq "Employee" and (emails co "example.com" or emails.value co "example.org")`,
		bson.M{"$and": []interface{}{
			bson.M{"$and": []interface{}{
				bson.M{"userType": bson.M{"$eq": "Employee"}},
				bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}},
			bson.M{"$or": []interface{}{
				bson.M{"$elemMatch": bson.M{"$and": []interface{}{
					bson.M{"emails.value": bson.M{"$regex": bson.RegEx{Pattern: "example\\.com", Options: "i"}}},
					bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}}},
				bson.M{"$elemMatch": bson.M{"$and": []interface{}{
					bson.M{"emails.value": bson.M{"$regex": bson.RegEx{Pattern: "example\\.org", Options: "i"}}},
					bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}}}}}},
			"meta.resourceType": "User",
		},
	},
	{
		`userType eq "Employee" and emails.type eq "work"`,
		bson.M{"meta.resourceType": "User", "$and": []interface{}{
			bson.M{"$and": []interface{}{
				bson.M{"userType": bson.M{"$eq": "Employee"}},
				bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}},
			bson.M{"$elemMatch": bson.M{"$and": []interface{}{
				bson.M{"emails.type": bson.M{"$eq": "work"}},
				bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}}}},
		},
	},
	{
		`userType eq "Employee" and (emails.type eq "work")`,
		bson.M{"meta.resourceType": "User", "$and": []interface{}{
			bson.M{"$and": []interface{}{
				bson.M{"userType": bson.M{"$eq": "Employee"}},
				bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}},
			bson.M{"$elemMatch": bson.M{"$and": []interface{}{
				bson.M{"emails.type": bson.M{"$eq": "work"}},
				bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}}}},
		},
	},
	{
		`emails.type eq "work"`,
		bson.M{"meta.resourceType": "User", "$elemMatch": bson.M{"$and": []interface{}{
			bson.M{"emails.type": bson.M{"$eq": "work"}},
			bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}},
		},
	},
	{
		`userName eq "bjensen" and name.familyName sw "J"`,
		bson.M{"meta.resourceType": "User", "$and": []interface{}{
			bson.M{"$and": []interface{}{
				bson.M{"userName": bson.M{"$eq": "bjensen"}},
				bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}},
			bson.M{"$and": []interface{}{
				bson.M{"name.familyName": bson.M{"$regex": bson.RegEx{Pattern: "^J", Options: "i"}}},
				bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}}},
		},
	},
	{
		`userName sw "J"`,
		bson.M{"$and": []interface{}{
			bson.M{"userName": bson.M{"$regex": bson.RegEx{Pattern: "^J", Options: "i"}}},
			bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}},
			"meta.resourceType": "User",
		},
	},
	{
		`emails co "example.com"`,
		bson.M{"$elemMatch": bson.M{"$and": []interface{}{
			bson.M{"emails.value": bson.M{"$regex": bson.RegEx{Pattern: "example\\.com", Options: "i"}}},
			bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}},
			"meta.resourceType": "User",
		},
	},
	{
		`emails.type co "work"`,
		bson.M{"$elemMatch": bson.M{"$and": []interface{}{
			bson.M{"emails.type": bson.M{"$regex": bson.RegEx{Pattern: "work", Options: "i"}}},
			bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}},
			"meta.resourceType": "User",
		},
	},
	{
		`emails.type ne true`,
		bson.M{"$elemMatch": bson.M{"$and": []interface{}{
			bson.M{"emails.type": bson.M{"$ne": true}},
			bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}},
			"meta.resourceType": "User",
		},
	},
	{
		`userType ne "Employee" and not (emails co "example.com" or emails.value co "example.org")`,
		bson.M{"$and": []interface{}{
			bson.M{"$and": []interface{}{
				bson.M{"userType": bson.M{"$ne": "Employee"}},
				bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}},
			bson.M{"$nor": []interface{}{
				bson.M{"$or": []interface{}{
					bson.M{"$elemMatch": bson.M{"$and": []interface{}{
						bson.M{"emails.value": bson.M{"$regex": bson.RegEx{Pattern: "example\\.com", Options: "i"}}},
						bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}}},
					bson.M{"$elemMatch": bson.M{"$and": []interface{}{
						bson.M{"emails.value": bson.M{"$regex": bson.RegEx{Pattern: "example\\.org", Options: "i"}}},
						bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}}}}}}}},
			"meta.resourceType": "User",
		},
	},
	{
		`userName eq "bjensen" and name.familyName sw "J"`,
		bson.M{"$and": []interface{}{
			bson.M{"$and": []interface{}{
				bson.M{"userName": bson.M{"$eq": "bjensen"}},
				bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}},
			bson.M{"$and": []interface{}{
				bson.M{"name.familyName": bson.M{"$regex": bson.RegEx{Pattern: "^J", Options: "i"}}},
				bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}}},
			"meta.resourceType": "User",
		},
	},
	{
		`userType ne "Employee" and not (emails co "example.com" or emails.value co "example.org")`,
		bson.M{"$and": []interface{}{
			bson.M{"$and": []interface{}{
				bson.M{"userType": bson.M{"$ne": "Employee"}},
				bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}},
			bson.M{"$nor": []interface{}{
				bson.M{"$or": []interface{}{
					bson.M{"$elemMatch": bson.M{"$and": []interface{}{
						bson.M{"emails.value": bson.M{"$regex": bson.RegEx{Pattern: "example\\.com", Options: "i"}}},
						bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}}},
					bson.M{"$elemMatch": bson.M{"$and": []interface{}{
						bson.M{"emails.value": bson.M{"$regex": bson.RegEx{Pattern: "example\\.org", Options: "i"}}},
						bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}}}}}}}},
			"meta.resourceType": "User",
		},
	},
	{
		`userName eq "bjensen"`,
		bson.M{"$and": []interface{}{
			bson.M{"userName": bson.M{"$eq": "bjensen"}},
			bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}},
			"meta.resourceType": "User",
		},
	},
	{
		`meta.lastModified gt "2011-05-13T04:42:34Z"`,
		bson.M{"$and": []interface{}{
			bson.M{"meta.lastModified": bson.M{"$gt": "2011-05-13T04:42:34Z"}},
			bson.M{"_uri": bson.M{"$eq": ""}}},
			"meta.resourceType": "User",
		},
	},
	{
		`meta.lastModified ge "2011-05-13T04:42:34Z"`,
		bson.M{"$and": []interface{}{
			bson.M{"meta.lastModified": bson.M{"$gte": "2011-05-13T04:42:34Z"}},
			bson.M{"_uri": bson.M{"$eq": ""}}},
			"meta.resourceType": "User",
		},
	},
	{
		`meta.lastModified lt "2011-05-13T04:42:34Z"`,
		bson.M{"$and": []interface{}{
			bson.M{"meta.lastModified": bson.M{"$lt": "2011-05-13T04:42:34Z"}},
			bson.M{"_uri": bson.M{"$eq": ""}}},
			"meta.resourceType": "User",
		},
	},
	{
		`meta.lastModified le "2011-05-13T04:42:34Z"`,
		bson.M{"$and": []interface{}{
			bson.M{"meta.lastModified": bson.M{"$lte": "2011-05-13T04:42:34Z"}},
			bson.M{"_uri": bson.M{"$eq": ""}}},
			"meta.resourceType": "User",
		},
	},
	{
		`name.familyName co "O'Malley"`,
		bson.M{"$and": []interface{}{
			bson.M{"name.familyName": bson.M{"$regex": bson.RegEx{Pattern: "O'Malley", Options: "i"}}},
			bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}},
			"meta.resourceType": "User",
		},
	},
	{
		`not (userName eq "strings")`,
		bson.M{"$nor": []interface{}{
			bson.M{"$and": []interface{}{
				bson.M{"userName": bson.M{"$eq": "strings"}},
				bson.M{"_uri": bson.M{"$eq": "urn:ietf:params:scim:schemas:core:2.0:User"}}}}},
			"meta.resourceType": "User",
		},
	},
}

func ierror(index int) string {
	return "Test case num. " + strconv.Itoa(index+1)
}

func herror(index int, test mongoCase) string {
	return ierror(index) + ", input `" + test.filter + "`"
}

func TestConvertToMongoQuery(t *testing.T) {
	dat, err := ioutil.ReadFile("../../internal/testdata/user.json")
	require.NotNil(t, dat)
	require.Nil(t, err)

	res := core.ResourceType{}
	json.Unmarshal(dat, &res)

	// for ii, tt := range mongoTests {
	// 	t.Run("AllAttributesUnknown", func(t *testing.T) {
	// 		ft, err := filter.CompileString(tt.filter)
	// 		require.NoError(t, err, herror(ii, tt))

	// 		m, err := convertToMongoQuery(&res, ft)
	// 		require.NoError(t, err, herror(ii, tt))

	// 		require.Equal(t, tt.query1, m, herror(ii, tt))
	// 	})
	// }

	// Now we load within repositories the resource type and the schema
	resTypeRepo := core.GetResourceTypeRepository()
	_, err = resTypeRepo.Add("../../internal/testdata/user.json")
	require.NoError(t, err)

	schemaRepo := core.GetSchemaRepository()
	_, err = schemaRepo.Add("../../internal/testdata/user_schema.json")
	require.NoError(t, err)

	for ii, tt := range mongoTests {
		t.Run("ValidAttributesKnown", func(t *testing.T) {
			ft, err := filter.CompileString(tt.filter)
			require.NoError(t, err, herror(ii, tt))

			m, err := convertToMongoQuery(&res, ft)
			require.NoError(t, err, herror(ii, tt))

			require.Equal(t, tt.query2, m, herror(ii, tt))
		})
	}
}
