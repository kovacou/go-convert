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
		output = v
	case []byte:
		output = string(v)
	case int:
		output = strconv.Itoa(v)
	case int8:
		output = strconv.Itoa(int(v))
	case int16:
		output = strconv.Itoa(int(v))
	case int32:
		output = strconv.Itoa(int(v))
	case int64:
		output = strconv.FormatInt(v, 10)
	case uint:
		output = strconv.FormatUint(uint64(v), 10)
	case uint8:
		output = strconv.FormatUint(uint64(v), 10)
	case uint16:
		output = strconv.FormatUint(uint64(v), 10)
	case uint32:
		output = strconv.FormatUint(uint64(v), 10)
	case uint64:
		output = strconv.FormatUint(v, 10)
	case float32:
		output = strconv.FormatFloat(float64(v), 'g', -1, 64)
	case float64:
		output = strconv.FormatFloat(v, 'g', -1, 64)
	case nil:
		output = ""
	case bool:
		output = strconv.FormatBool(v)
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
