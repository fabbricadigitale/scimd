package resource

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/fabbricadigitale/scimd/schemas/datatype"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalResource(t *testing.T) {
	resTypeRepo := core.GetResourceTypeRepository()
	if _, err := resTypeRepo.Add("../core/testdata/user.json"); err != nil {
		t.Log(err)
		t.Fail()
	}

	schemaRepo := core.GetSchemaRepository()
	if _, err := schemaRepo.Add("../core/testdata/user_schema.json"); err != nil {
		t.Log(err)
		t.Fail()
	}
	if _, err := schemaRepo.Add("../core/testdata/enterprise_user_schema.json"); err != nil {
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

	baseAttr := *res.Values("urn:ietf:params:scim:schemas:core:2.0:User")
	assert.Equal(t, true, !datatype.IsNull(baseAttr))

	extAttr := *res.Values("urn:ietf:params:scim:schemas:extension:enterprise:2.0:User")
	assert.Equal(t, true, !datatype.IsNull(extAttr))

	attrEqualities := []struct {
		value interface{}
		field interface{}
	}{
		{datatype.String("bjensen@example.com"), baseAttr["userName"]},
		{datatype.String("Babs Jensen"), baseAttr["displayName"]},
		{datatype.Boolean(true), baseAttr["active"]},
		{datatype.String("Ms. Barbara J Jensen, III"), baseAttr["name"].(datatype.Complex)["formatted"]},

		{datatype.String("701984"), extAttr["employeeNumber"]},
		{datatype.String("4130"), extAttr["costCenter"]},
	}

	for _, row := range attrEqualities {
		assert.Equal(t, row.value, row.field)
	}

}

func TestMarshalResource(t *testing.T) {
	resTypeRepo := core.GetResourceTypeRepository()
	if _, err := resTypeRepo.Add("../core/testdata/user.json"); err != nil {
		t.Log(err)
		t.Fail()
	}

	schemaRepo := core.GetSchemaRepository()
	if _, err := schemaRepo.Add("../core/testdata/user_schema.json"); err != nil {
		t.Log(err)
		t.Fail()
	}
	if _, err := schemaRepo.Add("../core/testdata/enterprise_user_schema.json"); err != nil {
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
