package resource

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/fabbricadigitale/scimd/schemas"
	"github.com/fabbricadigitale/scimd/schemas/core"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalResource(t *testing.T) {
	resTypeRepo := schemas.GetResourceTypeRepository()
	if _, err := resTypeRepo.Add("../testdata/user.json"); err != nil {
		t.Log(err)
		t.Fail()
	}

	schemaRepo := schemas.GetSchemaRepository()
	if _, err := schemaRepo.Add("../testdata/user_schema.json"); err != nil {
		t.Log(err)
		t.Fail()
	}
	if _, err := schemaRepo.Add("../testdata/enterprise_user_schema.json"); err != nil {
		t.Log(err)
		t.Fail()
	}

	// Non-normative of SCIM user resource type [https://tools.ietf.org/html/rfc7643#section-8.2]
	dat, err := ioutil.ReadFile("testdata/enterprise_user_resource.json")

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	require.NotNil(t, dat)
	require.Nil(t, err)

	res := Resource{}
	err = json.Unmarshal(dat, &res)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	equalities := []struct {
		value interface{}
		field interface{}
	}{
		{"2819c223-7f76-453a-919d-413861904646", res.ID},
		{"2819c223-7f76-453a-919d-413861904646", res.Common.ID},
		{[]string{"urn:ietf:params:scim:schemas:core:2.0:User", "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"}, res.Schemas},
		{[]string{"urn:ietf:params:scim:schemas:core:2.0:User", "urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"}, res.Common.Schemas},
		{"https://example.com/v2/Users/2819c223-7f76-453a-919d-413861904646", res.Meta.Location},
		{"https://example.com/v2/Users/2819c223-7f76-453a-919d-413861904646", res.Common.Meta.Location},
	}

	for _, row := range equalities {
		assert.Equal(t, row.value, row.field)
	}

	baseAttr := *res.GetValues("urn:ietf:params:scim:schemas:core:2.0:User")
	assert.Equal(t, true, !core.IsNull(baseAttr))

	extAttr := *res.GetValues("urn:ietf:params:scim:schemas:extension:enterprise:2.0:User")
	assert.Equal(t, true, !core.IsNull(extAttr))

	attrEqualities := []struct {
		value interface{}
		field interface{}
	}{
		{core.String("bjensen@example.com"), baseAttr["userName"]},
		{core.String("Babs Jensen"), baseAttr["displayName"]},
		{core.Boolean(true), baseAttr["active"]},
		{core.String("Ms. Barbara J Jensen, III"), baseAttr["name"].(core.Complex)["formatted"]},

		{core.String("701984"), extAttr["employeeNumber"]},
		{core.String("4130"), extAttr["costCenter"]},
	}

	for _, row := range attrEqualities {
		assert.Equal(t, row.value, row.field)
	}

}

func TestMarshalResource(t *testing.T) {
	resTypeRepo := schemas.GetResourceTypeRepository()
	if _, err := resTypeRepo.Add("../testdata/user.json"); err != nil {
		t.Log(err)
		t.Fail()
	}

	schemaRepo := schemas.GetSchemaRepository()
	if _, err := schemaRepo.Add("../testdata/user_schema.json"); err != nil {
		t.Log(err)
		t.Fail()
	}
	if _, err := schemaRepo.Add("../testdata/enterprise_user_schema.json"); err != nil {
		t.Log(err)
		t.Fail()
	}

	// Non-normative of SCIM user resource type [https://tools.ietf.org/html/rfc7643#section-8.2]
	dat, err := ioutil.ReadFile("testdata/enterprise_user_resource.json")

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	require.NotNil(t, dat)
	require.Nil(t, err)

	res := Resource{}
	err = json.Unmarshal(dat, &res)

	b, err := json.Marshal(&res)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	assert.NotEqual(t, 0, len(b))

	t.Logf("%s", string(b[:]))

}
