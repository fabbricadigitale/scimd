package mold

import (
	"context"
	"fmt"
	"math"
	"reflect"
	"strconv"

	m "gopkg.in/go-playground/mold.v2"
)

func max(ctx context.Context, t *m.Transformer, v reflect.Value, param string) (err error) {
	switch v.Interface().(type) {
	case int, int8, int16, int32, int64:
		if threshold, err := strconv.ParseInt(param, 10, 64); err == nil {
			if v.Int() > threshold {
				v.SetInt(threshold)
			}
		}
	case uint, uint8, uint16, uint32, uint64:
		if threshold, err := strconv.ParseUint(param, 10, 64); err == nil {
			if v.Uint() > threshold {
				v.SetUint(threshold)
			}
		}
	case float32, float64:
		if threshold, err := strconv.ParseFloat(param, 64); err == nil {
			v.SetFloat(math.Min(v.Float(), threshold))
		}
	default:
		err = &Error{
			message: fmt.Sprintf("Bad field type %T", v.Interface()),
		}
	}

	return
}
