package mongo

import (
	"reflect"

	"github.com/globalsign/mgo/bson"
)

type document bson.M

func (d *document) escapeKeys() {
	rv := reflect.ValueOf(d)
	apply(&rv, keyEscape, false)
}

func (d *document) unescapeKeys() {
	rv := reflect.ValueOf(d)
	apply(&rv, keyUnescape, true)
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

func zeroLen(rv reflect.Value) bool {
	rk := rv.Kind()
	if rk == reflect.Interface || rk == reflect.Ptr {
		rv = rv.Elem()
		rk = rv.Kind()
	}
	return (rk == reflect.Map || rk == reflect.Slice || rk == reflect.Array) && rv.Len() == 0
}

func deRefDeep(rv reflect.Value) reflect.Value {
	switch rv.Kind() {
	case reflect.Interface:
		return deRefDeep(rv.Elem())
	}
	return rv
}

func apply(rv *reflect.Value, f func(string) string, pruneEmpty bool) {

	switch rv.Kind() {

	case reflect.Array, reflect.Slice:
		l := rv.Len()

		newRv := reflect.MakeSlice(rv.Type(), 0, 0)
		for i := 0; i < l; i++ {
			rvv := deRefDeep(rv.Index(i))
			if pruneEmpty && zeroLen(rvv) {
				continue
			}
			apply(&rvv, f, pruneEmpty)
			newRv = reflect.Append(newRv, rvv)
		}
		(*rv) = newRv

	case reflect.Map:
		if rv.Type().Key().Kind() != reflect.String {
			panic("mongo: expecting map[string]bson.M, got map[" + rv.Type().Key().Kind().String() + "]" + rv.Elem().Type().String())
		}

		for _, k := range rv.MapKeys() {
			// unwind value
			rvv := deRefDeep(rv.MapIndex(k))

			if pruneEmpty && zeroLen(rvv) {
				rv.SetMapIndex(k, reflect.Value{}) // delete entry
				continue
			}

			apply(&rvv, f, pruneEmpty)

			// apply f() to key
			key := k.String()
			if newKey := f(key); key != newKey {
				// delete
				rv.SetMapIndex(k, reflect.Value{})
				// unescape key
				k = reflect.ValueOf(newKey)
			}

			// k or rvv may have been changed
			rv.SetMapIndex(k, rvv)
		}

	case reflect.Ptr:
		el := rv.Elem()
		apply(&el, f, pruneEmpty)
	}
}
