package convert

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// SnakeCase
func SnakeCase(in string) string {
	runes := []rune(in)
	b := strings.Builder{}
	b.Grow(len(in))

	for i, v := range runes {
		if i+1 < len(in) && unicode.IsLower(v) && unicode.IsUpper(runes[i+1]) {
			b.WriteRune(v)
			b.WriteRune('_')
		} else {
			b.WriteRune(unicode.ToLower(v))
		}
	}
	return b.String()
}

// StringE cast input to string
func StringE(input interface{}) (output string, err error) {
	switch v := input.(type) {
	case string:
		output = string(v)
		return
	case []byte:
		output = string(v)
		return
	case int:
		output = strconv.Itoa(v)
		return
	case int8:
		output = strconv.Itoa(int(v))
		return
	case int16:
		output = strconv.Itoa(int(v))
		return
	case int32:
		output = strconv.Itoa(int(v))
		return
	case int64:
		output = strconv.FormatInt(v, 10)
		return
	case uint:
		output = strconv.FormatUint(uint64(v), 10)
		return
	case uint8:
		output = strconv.FormatUint(uint64(v), 10)
		return
	case uint16:
		output = strconv.FormatUint(uint64(v), 10)
		return
	case uint32:
		output = strconv.FormatUint(uint64(v), 10)
		return
	case uint64:
		output = strconv.FormatUint(uint64(v), 10)
		return
	case float32:
		output = strconv.FormatFloat(float64(v), 'g', -1, 64)
		return
	case float64:
		output = strconv.FormatFloat(v, 'g', -1, 64)
		return
	case nil:
		output = ""
		return
	case bool:
		output = strconv.FormatBool(v)
		return
	default:
		err = errors.New("could not convert to string")
	}
	return
}

// String cast input to string and return default value if error
func String(input interface{}) string {
	o, _ := StringE(input)
	return o
}
