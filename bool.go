// Copyright © 2021 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package convert

import "errors"

// BoolE cast input to bool
func BoolE(input any) (out bool, err error) {
	switch v := input.(type) {
	case bool:
		out = v
	case int:
		out = v > 0
		return
	case int8:
		out = v > 0
		return
	case int16:
		out = v > 0
		return
	case int32:
		out = v > 0
		return
	case int64:
		out = v > 0
		return
	case uint:
		out = v > 0
		return
	case uint8:
		out = v > 0
		return
	case uint16:
		out = v > 0
		return
	case uint32:
		out = v > 0
		return
	case uint64:
		out = v > 0
		return
	case float32:
		out = v > 0
		return
	case float64:
		out = v > 0
		return
	case string:
		out = v == "true" || v == "1"
		return
	case []byte:
		out = Bool(string(v))
		return
	case nil:
		out = false
		return

	default:
		err = errors.New("could not convert to bool")
	}

	return
}

// Bool cast v to bool and return default value if error.
func Bool(v any) bool {
	out, _ := BoolE(v)
	return out
}
