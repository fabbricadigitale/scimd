package mold

import (
	"context"
	"fmt"
	"reflect"

	"github.com/leodido/go-urn"

	m "gopkg.in/go-playground/mold.v2"
)

func normurn(ctx context.Context, t *m.Transformer, v reflect.Value, param string) (e error) {
	switch v.Interface().(type) {
	case string:
		val := v.String()
		u, ok := urn.Parse(val)
		if !ok {
			e = &Error{
				message: fmt.Sprintf("Not a valid URN %s", val),
			}
		} else {
			v.SetString(u.Normalize().String())
		}
	default:
		e = &Error{
			message: fmt.Sprintf("Bad field type %T", v.Interface()),
		}
	}
}
