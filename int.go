package convert

import (
	"errors"
	"strconv"
)

// IntE cast input to int
func IntE(input interface{}) (output int, err error) {
	switch v := input.(type) {
	case string:
		var o int64
		o, err = strconv.ParseInt(v, 0, 0)
		output = int(o)
		return
	case []byte:
		var o int64
		o, err = strconv.ParseInt(string(v), 0, 0)
		output = int(o)
		return
	case int:
		output = int(v)
		return
	case int8:
		output = int(v)
		return
	case int16:
		output = int(v)
		return
	case int32:
		output = int(v)
		return
	case int64:
		output = int(v)
		return
	case uint:
		output = int(v)
		return
	case uint8:
		output = int(v)
		return
	case uint16:
		output = int(v)
		return
	case uint32:
		output = int(v)
		return
	case uint64:
		output = int(v)
		return
	case float32:
		output = int(v)
		return
	case float64:
		output = int(v)
		return
	case bool:
		output = 0
		if v {
			output = 1
		}
		return
	case nil:
		output = 0
		return
	default:
		err = errors.New("could not convert to int")
	}
	return
}

// Int cast i to int and return default value if error.
func Int(i interface{}) int {
	o, _ := IntE(i)
	return o
}
