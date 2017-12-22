package validation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type isSetOk struct {
	First  int
	Second int `validate:"isset=First"`
}

type isSetChanField struct {
	First  chan int
	Second int `validate:"isset=First"`
}

type isSetFuncField struct {
	First  func()
	Second int `validate:"isset=First"`
}

type isSetMapField struct {
	First  map[string]string
	Second int `validate:"isset=First"`
}

type isSetPtrField struct {
	First  *string
	Second int `validate:"isset=First"`
}

type isSetInterfaceField struct {
	First  interface{}
	Second int `validate:"isset=First"`
}

type isSetSliceField struct {
	First  []int
	Second int `validate:"isset=First"`
}

type isSetEmpty struct {
	First  int
	Second int `validate:"isset="`
}

type isSetNotFound struct {
	First  string
	Second int `validate:"isset=f"`
}

func TestIsSet(t *testing.T) {
	x := isSetOk{}
	x.First = 1
	x.Second = 2

	err := Validator.Struct(x)
	require.NoError(t, err)

	// Field type chan
	// Not setted
	a := isSetChanField{}
	a.Second = 2

	err = Validator.Struct(a)
	require.Error(t, err)

	// Setted
	a.First = make(chan int)
	err = Validator.Struct(a)
	require.NoError(t, err)

	// Field type func
	// Not setted
	b := isSetFuncField{}
	b.Second = 2

	err = Validator.Struct(b)
	require.Error(t, err)

	// Setted
	b.First = func() {}
	err = Validator.Struct(b)
	require.NoError(t, err)

	// Field type map
	// Not setted
	c := isSetMapField{}
	c.Second = 2

	err = Validator.Struct(c)
	require.Error(t, err)

	// Setted
	c.First = make(map[string]string)
	err = Validator.Struct(c)
	require.NoError(t, err)

	// Field type pointer
	// Not setted
	d := isSetPtrField{}
	d.Second = 2

	err = Validator.Struct(d)
	require.Error(t, err)

	// Setted
	str := ""
	d.First = &(str)
	err = Validator.Struct(d)
	require.NoError(t, err)

	// Field type interface
	// Not setted
	e := isSetInterfaceField{}
	e.Second = 2

	err = Validator.Struct(e)
	require.Error(t, err)

	// Setted
	e.First = 1
	err = Validator.Struct(e)
	require.NoError(t, err)

	// Field type slice
	// Not setted
	f := isSetSliceField{}
	f.Second = 2

	err = Validator.Struct(f)
	require.Error(t, err)

	// Setted
	f.First = make([]int, 5)
	err = Validator.Struct(f)
	require.NoError(t, err)

	// Empty isset tag parameter
	y := isSetEmpty{}
	y.First = 1
	y.Second = 2

	require.PanicsWithValue(t, "Invalid tag composition", func() {
		Validator.Struct(y)
	})

	z := isSetNotFound{}
	z.First = "hello"
	z.Second = 2

	require.PanicsWithValue(t, "Field not found", func() {
		Validator.Struct(z)
	})

	// Parent isn't a struct
	w := "hello"
	require.PanicsWithValue(t, "Invalid parent type string: must be a struct", func() {
		Validator.Var(w, "isset=y")
	})

}
