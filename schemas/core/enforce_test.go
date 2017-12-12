package core

import (
	"testing"

	"github.com/fabbricadigitale/scimd/schemas/datatype"
	"github.com/stretchr/testify/require"
)

func TestEnforceAttribute(t *testing.T) {

	attr := &Attribute{
		Name:        "userName",
		MultiValued: false,
		Type:        "string",
	}

	var value interface{}
	value = "john"

	enforced, err := attr.Enforce(value)
	require.NoError(t, err)

	datatypeString := datatype.String(value.(string))

	require.IsType(t, enforced, datatypeString)
}

func TestEnforceAttributes(t *testing.T) {

	attrs := &Attributes{
		&Attribute{
			Name:        "userName",
			MultiValued: false,
			Type:        "string",
		},
		&Attribute{
			Name:        "emails",
			MultiValued: true,
			Type:        "complex",
			SubAttributes: Attributes{
				&Attribute{
					Name:        "value",
					MultiValued: false,
					Type:        "string",
				},
				&Attribute{
					Name:        "type",
					MultiValued: false,
					Type:        "string",
				},
			},
		},
	}

	values := make(map[string]interface{})
	var singleValue interface{}
	singleValue = "john"
	values["userName"] = singleValue

	multiValue := []map[string]interface{}{
		{
			"value": "something@tin.su",
			"type":  "work",
		},
	}
	values["emails"] = multiValue

	enforcedValues, err := attrs.Enforce(values)
	require.NoError(t, err)

	datatypeString := datatype.String(singleValue.(string))
	require.Implements(t, (*datatype.DataTyper)(nil), (*enforcedValues)["userName"])
	require.IsType(t, datatypeString, (*enforcedValues)["userName"])

	require.IsType(t, [](datatype.DataTyper)(nil), (*enforcedValues)["emails"])
	for _, e := range (*enforcedValues)["emails"].([]datatype.DataTyper) {
		require.IsType(t, &datatype.Complex{}, e)
	}
}
