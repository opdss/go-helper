package strings

import (
	"errors"
	"strings"
	"testing"
)

var str1 = "a=1;b=2;c=3=4;d=5"
var strS = "a:1;b:2;c:3=4;d:5"
var strQ = "a=1&b=2&c=3=4&d=5"

func TestKv2MapFunc(t *testing.T) {
	res := Kv2MapFunc(str1, ";", func(s string) (string, string, error) {
		arr := strings.SplitN(s, "=", 2)
		if len(arr) != 2 {
			return "", "", errors.New("err")
		}
		return arr[0], arr[1], nil
	})
	if res["c"] != "3=4" {
		t.Errorf("Kv2MapFunc(%s) = %+v error", str1, res)
	}
}

func TestQuery2MapString(t *testing.T) {
	res := Query2MapString(strQ)
	if res["c"] != "3=4" || res["d"] != "5" {
		t.Errorf("Kv2MapFunc(%s) = %+v", str1, res)
	}
}

func TestKv2MapString(t *testing.T) {
	res := Kv2MapString(strS)
	if res["c"] != "3=4" || res["d"] != "5" {
		t.Errorf("Kv2MapFunc(%s) = %+v", str1, res)
	}
}
