package convert

import (
	"errors"
	"strconv"
)

// Float64E casts input to float64
func Float64E(v interface{}) (o float64, err error) {
	switch val := v.(type) {
	case string:
		o, err = strconv.ParseFloat(val, 64)
	case []byte:
		o, err = strconv.ParseFloat(string(val), 64)
	case int:
		o = float64(val)
	case int8:
		o = float64(val)
	case int16:
		o = float64(val)
	case int32:
		o = float64(val)
	case int64:
		o = float64(val)
	case uint:
		o = float64(val)
	case uint8:
		o = float64(val)
	case uint16:
		o = float64(val)
	case uint32:
		o = float64(val)
	case uint64:
		o = float64(val)
	case float32:
		o = float64(val)
	case float64:
		o = val
	case bool:
		if o = 0; val {
			o = 1
		}
	case nil:
	default:
		err = errors.New("could not convert to float64")
	}
	return
}

// Float64 cast v to float64 and return default value if there is any error.
func Float64(v interface{}) float64 {
	o, _ := Float64E(v)
	return o
}
