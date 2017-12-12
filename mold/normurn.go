package mold

import (
	"context"
	"fmt"
	"reflect"

	urn "github.com/leodido/go-urn"
	m "gopkg.in/go-playground/mold.v2"
)

// InvalidUrnError represents an error occuring during the urn normalization
// in case it's an invalid urn
type InvalidUrnError struct {
	message string
}

// BadTypeError epresents an error occuring during the urn normalization
// in case the input it's an invalid type
type BadTypeError struct {
	message string
}

func (e InvalidUrnError) Error() string { return e.message }

func (e BadTypeError) Error() string { return e.message }

func normurn(ctx context.Context, t *m.Transformer, v reflect.Value, param string) (e error) {
	switch v.Interface().(type) {
	case string:
		val := v.String()
		u, ok := urn.Parse(val)
		if !ok {
			e = &InvalidUrnError{
				message: fmt.Sprintf("Not a valid URN %s", val),
			}
		} else {
			v.SetString(u.Normalize().String())
		}
	default:
		e = &BadTypeError{
			message: fmt.Sprintf("Bad field type %T", v.Interface()),
		}
	}

	return
}
