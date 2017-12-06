package mongo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/fabbricadigitale/scimd/api/filter"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/stretchr/testify/require"
)

var filters = []string{
	`emails[type eq "work" and value co "@example.com"]`,
	`emails[type eq "work" and value co "@example.com"] or ims[type eq "xmpp" and value co "@foo.com"]`,
	`userType eq "Employee" and emails[type eq "work" and value co "@example.com"]`,
	`title pr`,
	`emails[not (type sw null)]`,
	`userType eq "Employee" and (emails co "example.com" or emails.value co "example.org")`,
	`userType eq "Employee" and emails.type eq "work"`,
	`userType eq "Employee" and (emails.type eq "work")`,
	`emails.type eq "work"`,
	`userName eq "bjensen" and name.familyName sw "J"`,
	`not (userName.Child eq "strings")`,
	`userName sw "J"`,
	`emails co "example.com"`,
	`emails.type co "work"`,
	`emails.type ne true`,
	`userType ne "Employee" and not (emails co "example.com" or emails.value co "example.org")`,
	`userName eq "bjensen" and name.familyName sw "J"`,
	`userType ne "Employee" and not (emails co "example.com" or emails.value co "example.org")`,
	`userName eq "bjensen"`,
	`meta.lastModified gt "2011-05-13T04:42:34Z"`,
	`meta.lastModified ge "2011-05-13T04:42:34Z"`,
	`meta.lastModified lt "2011-05-13T04:42:34Z"`,
	`meta.lastModified le "2011-05-13T04:42:34Z"`,
	`name.familyName co "O'Malley"`,
	`not (userName eq "strings")`,
}

func TestConvertToMongoQuery(t *testing.T) {

	for i, str := range filters {

		resTypeRepo := core.GetResourceTypeRepository()
		if _, err := resTypeRepo.Add("../../schemas/core/testdata/user.json"); err != nil {
			t.Log(err)
			t.Fail()
		}

		schemaRepo := core.GetSchemaRepository()
		if _, err := schemaRepo.Add("../../schemas/core/testdata/user_schema.json"); err != nil {
			t.Log(err)
			t.Fail()
		}

		dat, err := ioutil.ReadFile("../../schemas/core/testdata/user.json")

		require.NotNil(t, dat)
		require.Nil(t, err)

		res := core.ResourceType{}
		json.Unmarshal(dat, &res)

		ft, err := filter.CompileString(str)
		if err != nil {
			t.Log(err)
		}
		m, err := convertToMongoQuery(&res, ft)
		if err != nil {
			t.Log(err)
		}
		fmt.Printf("%d ----------------\n", i)
		fmt.Printf("Filter %s\n", str)
		fmt.Printf("Converted: %+v\n\n\n", m)
	}

}

func TestCreate(t *testing.T) {
	resTypeRepo := core.GetResourceTypeRepository()
	if _, err := resTypeRepo.Add("../../schemas/core/testdata/user.json"); err != nil {
		t.Log(err)
		t.Fail()
	}

	schemaRepo := core.GetSchemaRepository()
	if _, err := schemaRepo.Add("../../schemas/core/testdata/user_schema.json"); err != nil {
		t.Log(err)
		t.Fail()
	}
	if _, err := schemaRepo.Add("../../schemas/core/testdata/enterprise_user_schema.json"); err != nil {
		t.Log(err)
		t.Fail()
	}

	// Non-normative of SCIM user resource type [https://tools.ietf.org/html/rfc7643#section-8.2]
	dat, err := ioutil.ReadFile("../../schemas/resource/testdata/enterprise_user_resource_1.json")

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	require.NotNil(t, dat)
	require.Nil(t, err)

	res := &resource.Resource{}
	err = json.Unmarshal(dat, res)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	adapter, err := New("mongodb://localhost:27017", "scimd", "resources")

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	err = adapter.Create(res)

	if err != nil {
		t.Log(err)
	}

	require.Nil(t, err)
}

// (TODO) > Test hydrateResource adapter method

// (TODO) > Test toResource adapter method

// (TODO) > Test Get adapter method

// (TODO) > Test Delete adapter method

// (TODO) > Test Update adapter method
