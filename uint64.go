// Copyright © 2021 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package convert

import (
	"errors"
	"strconv"
)

// Uint64E Cast input to uint64
func Uint64E(v interface{}) (o uint64, err error) {
	switch val := v.(type) {
	case string:
		o, err = strconv.ParseUint(val, 10, 64)
	case []byte:
		o, err = strconv.ParseUint(string(val), 10, 64)
	case int:
		if val < 0 {
			err = errors.New("could not convert negative value to uint64")
			return
		}
		o = uint64(val)
	case int8:
		if val < 0 {
			err = errors.New("could not convert negative value to uint64")
			return
		}
		o = uint64(val)
	case int16:
		if val < 0 {
			err = errors.New("could not convert negative value to uint64")
			return
		}
		o = uint64(val)
	case int32:
		if val < 0 {
			err = errors.New("could not convert negative value to uint64")
			return
		}
		o = uint64(val)
	case int64:
		if val < 0 {
			err = errors.New("could not convert negative value to uint64")
			return
		}
		o = uint64(val)
	case uint:
		o = uint64(val)
	case uint8:
		o = uint64(val)
	case uint16:
		o = uint64(val)
	case uint32:
		o = uint64(val)
	case uint64:
		o = val
	case float32:
		if val < 0 {
			err = errors.New("could not convert negative value to uint64")
			return
		}
		o = uint64(val)
	case float64:
		if val < 0 {
			err = errors.New("could not convert negative value to uint64")
			return
		}
		o = uint64(val)
	case bool:
		if o = 0; val {
			o = 1
		}
	case nil:
	default:
		err = errors.New("could not convert to uint64")
	}
	return
}

// Uint64 cast v to uint64 and return default value if there is any error.
func Uint64(v interface{}) uint64 {
	o, _ := Uint64E(v)
	return o
}
