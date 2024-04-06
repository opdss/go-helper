package rand

const Alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func StringN(n int) string {
	str := make([]byte, n)
	l := len(Alpha)
	s := []byte(Alpha)
	for i := 0; i < n; i++ {
		str[i] = s[Int(0, l)]
	}
	return string(str)
}
