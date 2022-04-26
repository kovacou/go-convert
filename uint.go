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
func UintE(v any) (o uint, err error) {
	switch val := v.(type) {
	case string:
		v, err := strconv.ParseUint(val, 0, 0)
		if err == nil {
			o = uint(v)
		}
	case []byte:
		v, err := strconv.ParseUint(string(val), 0, 0)
		if err == nil {
			o = uint(v)
		}
	case int:
		if val < 0 {
			err = errors.New("could not convert negative value to uint")
			return
		}
		o = uint(val)
	case int8:
		if val < 0 {
			err = errors.New("could not convert negative value to uint")
			return
		}
		o = uint(val)
	case int16:
		if val < 0 {
			err = errors.New("could not convert negative value to uint")
			return
		}
		o = uint(val)
	case int32:
		if val < 0 {
			err = errors.New("could not convert negative value to uint")
			return
		}
		o = uint(val)
	case int64:
		if val < 0 {
			err = errors.New("could not convert negative value to uint")
			return
		}
		o = uint(val)
	case uint:
		o = val
	case uint8:
		o = uint(val)
	case uint16:
		o = uint(val)
	case uint32:
		o = uint(val)
	case uint64:
		o = uint(val)
	case float32:
		if val < 0 {
			err = errors.New("could not convert negative value to uint")
			return
		}
		o = uint(val)
	case float64:
		if val < 0 {
			err = errors.New("could not convert negative value to uint")
			return
		}
		o = uint(val)
	case bool:
		if o = 0; val {
			o = 1
		}
	case nil:
	default:
		err = errors.New("could not convert to uint")
	}
	return
}

// Uint cast v to uint64 and return default value if there is any error.
func Uint(v any) uint {
	o, _ := UintE(v)
	return o
}
