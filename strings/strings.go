package strings

import (
	"errors"
	"strings"
)

// Kv2MapString 将 k1:v1;k2:v2 字符串转成map
func Kv2MapString(str string) map[string]string {
	return Kv2MapFunc(str, ";", func(s string) (string, string, error) {
		arr := strings.SplitN(s, ":", 2)
		if len(arr) != 2 {
			return "", "", errors.New("Kv2MapString error")
		}
		return arr[0], arr[1], nil
	})
}

// Query2MapString 将 k1=v1&k2=v2 字符串转成map
func Query2MapString(str string) map[string]string {
	return Kv2MapFunc(str, "&", func(s string) (string, string, error) {
		arr := strings.SplitN(s, "=", 2)
		if len(arr) != 2 {
			return "", "", errors.New("Query2MapString error")
		}
		return arr[0], arr[1], nil
	})
}

// Kv2MapFunc 将 k1:v1;k2:v2 字符串转成map
func Kv2MapFunc[K comparable, V any](str string, sep string, fn func(s string) (K, V, error)) map[K]V {
	res := make(map[K]V, 0)
	if str == "" {
		return res
	}
	r := strings.Split(str, sep)
	for _, val := range r {
		k, v, err := fn(val)
		if err == nil {
			res[k] = v
		}
	}
	return res
}
