package str

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Length returns number of characters
func Length(s string) int {
	return len([]rune(s))
}

// Before returns the string before the first occurrence of the substr string
func Before(s, substr string) string {
	if substr == "" {
		return s
	}
	i := strings.Index(s, substr)
	if i != -1 {
		return s[:i]
	}
	return s
}

// BeforeLast returns the string before the last occurrence of the substr string
func BeforeLast(s, substr string) string {
	if substr == "" {
		return s
	}
	i := strings.LastIndex(s, substr)
	if i != -1 {
		return s[:i]
	}
	return s
}

// After returns the string after the first occurrence of the substr string
func After(s, substr string) string {
	if substr == "" {
		return s
	}
	i := strings.Index(s, substr)
	if i != -1 {
		i = i + len(substr)
		return s[i:]
	}
	return s
}

// AfterLast returns the string after the last occurrence of the substr string
func AfterLast(s, substr string) string {
	if substr == "" {
		return s
	}
	i := strings.LastIndex(s, substr)
	if i != -1 {
		i = i + len(substr)
		return s[i:]
	}
	return s
}

func Index(s, substr string) int {
	return strings.Index(s,substr)
}

func Contians(s, substr string) bool {
	return strings.Contains(s, substr)
}

func StartsWith(s, substr string) bool {
	if substr != "" && Substr(s, 0, len([]rune(substr))) == substr {
		return true
	}
	return false
}

func EndsWith(s, substr string) bool {
	if Substr(s, -len([]rune(substr)), len(s)) == substr {
		return true
	}
	return false
}

// Substr returns a string of length length from the start position
func Substr(s string, start int, strlength ...int) string {
	charlist := []rune(s)
	l := len(charlist)
	length :=0
	end := 0

	if len(strlength)==0{
		length = l
	}else{
		length = strlength[0]
	}

	if start < 0 {
		start = l + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}

	if start > l {
		start = l
	}

	if end < 0 {
		end = 0
	}

	if end > l {
		end = l
	}

	return string(charlist[start:end])
}

// Char returns a char slice
func Char(str string) []string {
	c := make([]string, 0)
	for _, v := range str {
		c = append(c, string(v))
	}
	return c
}

func Escape(s string) string {
	str := strconv.Quote(s)
	str = strings.Replace(str, "'", "\\'", -1)
	strlist := []rune(str)
	l := len(strlist)
	return Substr(str, 1, l-2)
}

// String returns a string of any type
func String(iface interface{}) string {
	switch val := iface.(type) {
	case []byte:
		return string(val)
	}
	v := reflect.ValueOf(iface)
	switch v.Kind() {
	case reflect.Invalid:
		return ""
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', -1, 32)
	case reflect.Ptr:
		b, err := json.Marshal(v.Interface())
		if err != nil {
			return "nil"
		}
		return string(b)
	}
	return fmt.Sprintf("%v", iface)
}