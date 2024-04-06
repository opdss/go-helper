package rand

import (
	"fmt"
	"testing"
)

func TestSlices(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	fmt.Println(a[:5])
	fmt.Println(a[5:])
	r := SlicesN(a, 0)
	fmt.Println(r)
	r1 := SlicesN(a, 9)
	fmt.Println(r1)
	r2 := SlicesN(a, 10)
	fmt.Println(r2)
	r3 := SlicesN(a, 13)
	fmt.Println(r3)
}
