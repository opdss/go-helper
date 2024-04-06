package rand

// SlicesN 在一个切片内随机获取n个元素
func SlicesN[T any](m []T, n int) []T {
	_m := make([]T, len(m))
	copy(_m, m)
	if len(m) <= n {
		return _m
	}
	r := make([]T, n)
	for i := 0; i < n; i++ {
		idx := Int(0, len(_m))
		r[i] = _m[idx]
		_m = append(_m[:idx], _m[idx+1:]...)
	}
	return r
}
