package mongo

import (
	"reflect"

	"github.com/globalsign/mgo/bson"
)

type document bson.M

func (d *document) escapeKeys() {
	rv := reflect.ValueOf(d)
	apply(&rv, keyEscape)
}

func (d *document) unescapeKeys() {
	rv := reflect.ValueOf(d)
	apply(&rv, keyUnescape)
}

func (d *document) GetBSON() (interface{}, error) {
	d.escapeKeys()
	return bson.M(*d), nil
}

func (d *document) SetBSON(raw bson.Raw) error {
	var m bson.M
	if err := raw.Unmarshal(&m); err != nil {
		return err
	}

	(*d) = document(m)
	d.unescapeKeys()

	return nil
}

func apply(rv *reflect.Value, f func(string) string) {
	switch rv.Kind() {
	case reflect.Array, reflect.Slice:
		l := rv.Len()
		for i := 0; i < l; i++ {
			rvv := rv.Index(i)
			apply(&rvv, f)
		}
	case reflect.Map:
		if rv.Type().Key().Kind() != reflect.String {
			panic("mongo: expecting map[string]bson.M, got map[" + rv.Type().Key().Kind().String() + "]" + rv.Elem().Type().String())
		}
		for _, k := range rv.MapKeys() {
			// unwind value
			rvv := rv.MapIndex(k)
			apply(&rvv, f)

			// apply f() to key
			key := k.String()
			if newKey := f(key); key != newKey {
				// delete
				rv.SetMapIndex(k, reflect.Value{})
				// unescape key
				k = reflect.ValueOf(newKey)
				// set value to new key
				rv.SetMapIndex(k, rvv)
			}
		}
	case reflect.Interface, reflect.Ptr:
		el := rv.Elem()
		apply(&el, f)
	}
}
