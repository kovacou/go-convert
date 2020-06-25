package convert

import (
	"errors"
	"strconv"
)

// Uint64E Cast input to uint64
func Uint64E(input interface{}) (output uint64, err error) {
	switch v := input.(type) {
	case string:
		output, err = strconv.ParseUint(v, 10, 64)
		return
	case []byte:
		output, err = strconv.ParseUint(string(v), 10, 64)
		return
	case int:
		if v < 0 {
			err = errors.New("could not convert negative value to uint64")
			return
		}
		output = uint64(v)
		return
	case int8:
		if v < 0 {
			err = errors.New("could not convert negative value to uint64")
			return
		}
		output = uint64(v)
		return
	case int16:
		if v < 0 {
			err = errors.New("could not convert negative value to uint64")
			return
		}
		output = uint64(v)
		return
	case int32:
		if v < 0 {
			err = errors.New("could not convert negative value to uint64")
			return
		}
		output = uint64(v)
		return
	case int64:
		if v < 0 {
			err = errors.New("could not convert negative value to uint64")
			return
		}
		output = uint64(v)
		return
	case uint:
		output = uint64(v)
		return
	case uint8:
		output = uint64(v)
		return
	case uint16:
		output = uint64(v)
		return
	case uint32:
		output = uint64(v)
		return
	case uint64:
		output = v
		return
	case float32:
		if v < 0 {
			err = errors.New("could not convert negative value to uint64")
			return
		}
		output = uint64(v)
		return
	case float64:
		if v < 0 {
			err = errors.New("could not convert negative value to uint64")
			return
		}
		output = uint64(v)
		return
	case bool:
		output = uint64(0)
		if v {
			output = uint64(1)
		}
		return
	case nil:
		output = uint64(0)
		return
	default:
		err = errors.New("could not convert to uint")
	}
	return
}

// Uint64 cast input to uint64 and return default value if error
func Uint64(input interface{}) uint64 {
	o, _ := Uint64E(input)
	return o
}
