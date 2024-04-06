package files

import (
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	src := "./copy.go"
	dst := "./copy.go_bak"
	n, err := Copy(dst, src)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(n)
}
