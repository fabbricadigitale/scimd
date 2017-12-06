package filter

import (
	"testing"

	"github.com/fabbricadigitale/scimd/api/attr"
	"github.com/fabbricadigitale/scimd/schemas/core"
	"github.com/stretchr/testify/assert"
)

const (
	filter1  = `userType ne "Employee" and not (emails co "example.com" or emails.value co "example.org")`
	filter2  = `user eq "bjensen"`
	filter3  = `emails.type ne true`
	filter4  = `name.familyName co "O'Malley"`
	filter5  = `userName sw "J"`
	filter6  = `meta.lastModified gt "2011-05-13T04:42:34Z"`
	filter7  = `meta.lastModified ge "2011-05-13T04:42:34Z"`
	filter8  = `meta.lastModified lt "2011-05-13T04:42:34Z"`
	filter9  = `meta.lastModified le "2011-05-13T04:42:34Z"`
	filter10 = `emails[type eq "work" and value co "@example.com"] or ims[type eq "xmpp" and value co "@foo.com"]`
	filter11 = `userType eq "Employee" and (emails co "example.com" or emails.value co "example.org")`
	filter12 = `userType eq "Employee" and emails[type eq "work" and value co "@example.com"]`
	filter13 = `userType eq "Employee" and (emails.type eq "work")`
	filter14 = `title pr`
	filter15 = `not (userName eq "strings")`
	filter16 = `not (userName.Child eq "strings")`
	filter17 = `emails[not (type sw null)]`
	filter18 = `title pr and userType eq "Employee"`
	filter19 = `title pr or userType eq "Intern"`
)

func TestStringer(t *testing.T) {
	var f1 Filter = And{
		AttrExpr{attr.Path{Name: "userType"}, OpNotEqual, "Employee"},
		Not{
			Or{
				AttrExpr{attr.Path{Name: "emails"}, OpContains, "example.com"},
				AttrExpr{attr.Path{Name: "emails", Sub: "value"}, OpContains, "example.org"},
			},
		},
	}

	assert.Equal(t, filter1, f1.String())

	var f2 Filter = AttrExpr{attr.Path{Name: "user"}, OpEqual, "bjensen"}

	assert.Equal(t, filter2, f2.String())

	var f3 Filter = AttrExpr{attr.Path{Name: "emails", Sub: "type"}, OpNotEqual, true}

	assert.Equal(t, filter3, f3.String())

	var f4 Filter = AttrExpr{attr.Path{Name: "name", Sub: "familyName"}, OpContains, "O'Malley"}

	assert.Equal(t, filter4, f4.String())

	var f5 Filter = AttrExpr{attr.Path{Name: "userName"}, OpStartsWith, "J"}

	assert.Equal(t, filter5, f5.String())

	var f6 Filter = AttrExpr{attr.Path{Name: "meta", Sub: "lastModified"}, OpGreaterThan, "2011-05-13T04:42:34Z"}

	assert.Equal(t, filter6, f6.String())

	var f7 Filter = AttrExpr{attr.Path{Name: "meta", Sub: "lastModified"}, OpGreaterOrEqualThan, "2011-05-13T04:42:34Z"}

	assert.Equal(t, filter7, f7.String())

	var f8 Filter = AttrExpr{attr.Path{Name: "meta", Sub: "lastModified"}, OpLessThan, "2011-05-13T04:42:34Z"}

	assert.Equal(t, filter8, f8.String())

	var f9 Filter = AttrExpr{attr.Path{Name: "meta", Sub: "lastModified"}, OpLessOrEqualThan, "2011-05-13T04:42:34Z"}

	assert.Equal(t, filter9, f9.String())

	var f10 Filter = Or{
		ValuePath{
			attr.Path{Name: "emails"},
			ValueAnd{
				AttrExpr{attr.Path{Name: "type"}, OpEqual, "work"},
				AttrExpr{attr.Path{Name: "value"}, OpContains, "@example.com"},
			},
		}, ValuePath{
			attr.Path{Name: "ims"},
			ValueAnd{
				AttrExpr{attr.Path{Name: "type"}, OpEqual, "xmpp"},
				AttrExpr{attr.Path{Name: "value"}, OpContains, "@foo.com"},
			},
		},
	}

	assert.Equal(t, filter10, f10.String())

	var f11 Filter = And{
		AttrExpr{attr.Path{Name: "userType"}, OpEqual, "Employee"},
		Group{
			Or{
				AttrExpr{attr.Path{Name: "emails"}, OpContains, "example.com"},
				AttrExpr{attr.Path{Name: "emails", Sub: "value"}, OpContains, "example.org"},
			},
		},
	}

	assert.Equal(t, filter11, f11.String())

	var f12 Filter = And{
		AttrExpr{attr.Path{Name: "userType"}, OpEqual, "Employee"},
		ValuePath{
			attr.Path{Name: "emails"},
			ValueAnd{
				AttrExpr{attr.Path{Name: "type"}, OpEqual, "work"},
				AttrExpr{attr.Path{Name: "value"}, OpContains, "@example.com"},
			},
		},
	}

	assert.Equal(t, filter12, f12.String())

	var f13 Filter = And{
		AttrExpr{attr.Path{Name: "userType"}, OpEqual, "Employee"},
		Group{
			AttrExpr{attr.Path{Name: "emails", Sub: "type"}, OpEqual, "work"},
		},
	}

	assert.Equal(t, filter13, f13.String())

	var f14 Filter = AttrExpr{attr.Path{Name: "title"}, OpPresent, nil}

	assert.Equal(t, filter14, f14.String())

	var f15 Filter = Not{
		AttrExpr{attr.Path{Name: "userName"}, OpEqual, "strings"},
	}

	assert.Equal(t, filter15, f15.String())

	var f16 Filter = Not{
		AttrExpr{attr.Path{Name: "userName", Sub: "Child"}, OpEqual, "strings"},
	}

	assert.Equal(t, filter16, f16.String())

	var f17 Filter = ValuePath{
		attr.Path{Name: "emails"},
		ValueNot{
			AttrExpr{attr.Path{Name: "type"}, OpStartsWith, nil},
		},
	}

	assert.Equal(t, filter17, f17.String())

	var f18 Filter = And{
		AttrExpr{attr.Path{Name: "title"}, OpPresent, nil},
		AttrExpr{attr.Path{Name: "userType"}, OpEqual, "Employee"},
	}
	assert.Equal(t, filter18, f18.String())

	var f19 Filter = Or{
		AttrExpr{attr.Path{Name: "title"}, OpPresent, nil},
		AttrExpr{attr.Path{Name: "userType"}, OpEqual, "Intern"},
	}
	assert.Equal(t, filter19, f19.String())
}

func loadRt(t *testing.T) {
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
}
func TestNormalize(t *testing.T) {
	loadRt(t)
	rt := core.GetResourceTypeRepository().Get("User")

	//
	f12, _ := CompileString(filter12)
	nf12 := f12.Normalize(rt)

	assert.Equal(
		t,
		`urn:ietf:params:scim:schemas:core:2.0:User:userType eq "Employee" and (urn:ietf:params:scim:schemas:core:2.0:User:emails.type eq "work" and urn:ietf:params:scim:schemas:core:2.0:User:emails.value co "@example.com")`,
		nf12.String(),
	)

	assert.Equal(
		t,
		nf12.String(),
		nf12.Normalize(rt).String(),
	)

	//

	f17, _ := CompileString(filter17)
	nf17 := f17.Normalize(rt)

	assert.Equal(
		t,
		`(not (urn:ietf:params:scim:schemas:core:2.0:User:emails.type sw null))`,
		nf17.String(),
	)

	assert.Equal(
		t,
		nf17.String(),
		nf17.Normalize(rt).String(),
	)

	// (todo) add more testing cases
}
