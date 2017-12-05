package mongo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/fabbricadigitale/scimd/api/filter"
	"github.com/fabbricadigitale/scimd/schemas/core"
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
