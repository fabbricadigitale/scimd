package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/fabbricadigitale/scimd/api"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/resource"
	"github.com/stretchr/testify/require"
)

var filters = []string{
	`userName eq "bjensen@example.com"`,
	`emails[type eq "work" and value co "@example.com"]`,
	`userName eq "bjensen" and name.familyName sw "J"`,
	`userType ne "Employee" and not (emails co "example.com" or emails.value co "example.org")`,
	`emails.type ne true`,
	`name.familyName co "O'Malley"`,
	`userName sw "J"`,
	`meta.lastModified gt "2011-05-13T04:42:34Z"`,
	`meta.lastModified ge "2011-05-13T04:42:34Z"`,
	`meta.lastModified lt "2011-05-13T04:42:34Z"`,
	`meta.lastModified le "2011-05-13T04:42:34Z"`,
	`emails[type eq "work" and value co "@example.com"] or ims[type eq "xmpp" and value co "@foo.com"]`,
	`userType eq "Employee" and (emails co "example.com" or emails.value co "example.org")`,
	`userType eq "Employee" and emails[type eq "work" and value co "@example.com"]`,
	`userType eq "Employee" and (emails.type eq "work")`,
	`title pr`,
	`not (userName eq "strings")`,
	`not (userName.Child eq "strings")`,
}

func TestCreateRepository(t *testing.T) {

	var manager Manager
	_, err := manager.CreateAdapter("mongo", "mongodb://localhost:27017/test_db?maxPoolSize=100", "test_db", "resources")

	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func TestCreate(t *testing.T) {
	resTypeRepo := core.GetResourceTypeRepository()
	if _, err := resTypeRepo.Add("../../internal/testdata/user.json"); err != nil {
		t.Log(err)
		t.Fail()
	}

	schemaRepo := core.GetSchemaRepository()
	if _, err := schemaRepo.Add("../../internal/testdata/user_schema.json"); err != nil {
		t.Log(err)
		t.Fail()
	}
	if _, err := schemaRepo.Add("../../internal/testdata/enterprise_user_schema.json"); err != nil {
		t.Log(err)
		t.Fail()
	}

	// Non-normative of SCIM user resource type [https://tools.ietf.org/html/rfc7643#section-8.2]
	dat, err := ioutil.ReadFile("../schemas/resource/testdata/enterprise_user_resource.json")

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

	var manager Manager
	adapter, err := manager.CreateAdapter("mongo", "mongodb://localhost:27017", "test_db", "resources")

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

func TestCountMethod(t *testing.T) {

	var manager Manager
	adapter, err := manager.CreateAdapter("mongo", "mongodb://localhost:27017", "test_db", "resources")

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	for i, filter := range filters {

		var f api.Filter
		f = api.Filter(filter)

		count, err := adapter.Count(nil, &f)
		if err != nil {
			t.Log(err)
		}

		require.Nil(t, err)
		fmt.Printf("%d ----------------\n", i)
		fmt.Printf("Filter %s\n", filter)
		fmt.Printf("Count: %v\n\n\n", count)
	}
}

// (TODO) > Test hydrateResource adapter method

// (TODO) > Test toResource adapter method

// (TODO) > Test Get adapter method

// (TODO) > Test Delete adapter method

// (TODO) > Test Update adapter method
