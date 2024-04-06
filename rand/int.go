package rand

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func Int(min, max int) int {
	return rand.Intn(max-min) + min
}

// IntN 在指定范围内生成指定数量的随机数,返回的数可能会有重复
func IntN(n1, n2 int, n int) []int {
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = Int(n1, n2)
	}
	return r
}
