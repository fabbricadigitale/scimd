package mold

import (
	"context"
	"fmt"
	"math"
	"reflect"
	"strconv"

	m "gopkg.in/go-playground/mold.v2"
)

func max(ctx context.Context, t *m.Transformer, v reflect.Value, param string) (e error) {
	switch v.Interface().(type) {
	case int, int8, int16, int32, int64:
		threshold, err := strconv.ParseInt(param, 10, 64)
		if err == nil {
			if v.Int() > threshold {
				v.SetInt(threshold)
			}
		}
		e = err
	case uint, uint8, uint16, uint32, uint64:
		threshold, err := strconv.ParseUint(param, 10, 64)
		if err == nil {
			if v.Uint() > threshold {
				v.SetUint(threshold)
			}
		}
		e = err
	case float32, float64:
		threshold, err := strconv.ParseFloat(param, 64) 
			if err == nil {
			v.SetFloat(math.Min(v.Float(), threshold))
			}
			e = err
	default:
		e = &Error{
			message: fmt.Sprintf("Bad field type %T", v.Interface()),
		}
	}

	return
}
