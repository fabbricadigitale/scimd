package messages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/fabbricadigitale/scimd/validation"
	"github.com/stretchr/testify/require"
)

func TestPatchResource(t *testing.T) {

	// Unmarshal
	dat, err := ioutil.ReadFile("testdata/patch.json")

	require.NotNil(t, dat)
	require.Nil(t, err)

	patch := PatchOp{}
	json.Unmarshal(dat, &patch)

	require.Equal(t, "urn:ietf:params:scim:api:messages:2.0:PatchOp", patch.Schemas[0])
	require.Equal(t, "add", patch.Operations[0].Op)
	require.Equal(t, "members", patch.Operations[0].Path)
	require.Equal(t, "Babs Jensen", patch.Operations[0].Value[0].Display)
	require.Equal(t, "https://example.com/v2/Users/2819c223", patch.Operations[0].Value[0].Ref)
	require.Equal(t, "2819c223-7f76", patch.Operations[0].Value[0].Value)
	require.Equal(t, "home", patch.Operations[0].Value[0].Type)

	// Marshal
	p := PatchOp{}
	p.Operations = make([]*Operation, 1)
	p.Schemas = []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"}
	value := []*Value{
		{
			Display: "Babs Jensen",
			Ref:     "https://example.com/v2/Users/2819c223",
			Value:   "2819c223-7f76",
			Type:    "home",
		},
	}
	p.Operations[0] = &Operation{
		Op:    "add",
		Path:  "members",
		Value: value,
	}

	b, err2 := json.Marshal(p)
	require.NotNil(t, b)
	require.Nil(t, err2)
	require.JSONEq(t, `{
		"schemas": ["urn:ietf:params:scim:api:messages:2.0:PatchOp"],
		"Operations": [{
			"op": "add",
			"path": "members",
			"value": [{
				"display": "Babs Jensen",
				"$ref": "https://example.com/v2/Users/2819c223",
				"value": "2819c223-7f76",
				"type": "home"
			}]
		}]
		}`, string(b))
	fmt.Println(string(b))
}

func TestPatchOpValid(t *testing.T) {
	var err error
	value := []*Value{
		{
			Display: "Babs Jensen",
			Ref:     "https://example.com/v2/Users/2819c223",
			Value:   "2819c223-7f76",
			Type:    "home",
		},
	}

	// // Right PatchOp validation tags
	schema := []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"}
	err = validation.Validator.Var(schema, "eq=1,dive,eq=urn:ietf:params:scim:api:messages:2.0:PatchOp")
	require.NoError(t, err)

	// Add: empty path is OK; Value field is required
	op := Operation{}
	op.Op = "add"
	op.Value = value
	err = validation.Validator.Struct(op)
	require.NoError(t, err)

	// Add: if path, must be an attrpath
	op = Operation{}
	op.Op = "add"
	op.Path = "name.firstName"
	op.Value = value
	err = validation.Validator.Struct(op)
	require.NoError(t, err)

	// (negative) Add: if path, must be an attrpath
	op = Operation{}
	op.Op = "add"
	op.Path = "#######" // wrong path
	op.Value = value
	err = validation.Validator.Struct(op)
	require.Error(t, err)

	// Replace: empty path is OK; Value field is required
	op = Operation{}
	op.Op = "replace"
	op.Value = value
	err = validation.Validator.Struct(op)
	require.NoError(t, err)

	// Replace: if path, must be an attrpath
	op = Operation{}
	op.Op = "replace"
	op.Path = "name.firstName"
	op.Value = value
	err = validation.Validator.Struct(op)
	require.NoError(t, err)

	// (negative) Replace: if path, must be an attrpath
	op = Operation{}
	op.Op = "replace"
	op.Path = "#######" // wrong path
	op.Value = value
	err = validation.Validator.Struct(op)
	require.Error(t, err)

	// Remove: Path cannot be empty; Value field is optional
	op = Operation{}
	op.Op = "remove"
	err = validation.Validator.Struct(op)
	require.Error(t, err)

	// Remove: Path must be an attrpath (not empty)
	op = Operation{}
	op.Op = "remove"
	op.Path = "name.firstName"
	op.Value = value
	err = validation.Validator.Struct(op)
	require.NoError(t, err)

	// Invalid value for Op field
	op = Operation{}
	op.Op = "abc"
	op.Path = "name.firstName"
	op.Value = value
	err = validation.Validator.Struct(op)
	require.Error(t, err)

	// Empty value for Op field
	op = Operation{}
	op.Path = "name.firstName"
	op.Value = value
	err = validation.Validator.Struct(op)
	require.Error(t, err)

	// Test Value struct: Value field required
	val := Value{
		Display: "Babs Jensen",
		Ref:     "https://example.com/v2/Users/2819c223",
		Value:   "2819c223-7f76",
		Type:    "home",
	}
	err = validation.Validator.Struct(val)
	require.NoError(t, err)

	// Value field empty
	val.Value = ""
	err = validation.Validator.Struct(val)
	require.Error(t, err)
}
