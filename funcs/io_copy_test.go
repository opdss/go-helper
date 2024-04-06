package funcs

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/opdss/go-helper/rand"
	"io"
	"testing"
	"time"
)

type src struct {
	s []byte
	c int
}

func (s *src) Read(b []byte) (int, error) {
	l := len(b)
	l2 := len(s.s)
	n := 0
	for i := 0; i < l; i++ {
		if s.c == l2 {
			return n, io.EOF
		}
		time.Sleep(time.Microsecond)
		b[i] = s.s[s.c]
		s.c++
		n++
	}
	return n, nil
}

func TestCopyWithContext(t *testing.T) {
	ss := []byte(rand.StringN(2000))
	s := &src{s: ss}
	d := &bytes.Buffer{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	n, err := CopyWithContext(ctx, d, s)
	if err != nil && !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("TestCopyWithContext error:%v", err)
		return
	}
	//n, err := CopyWithContext(ctx, d, s)
	//n, err := io.Copy(d, s)
	fmt.Print(n, err, d.Len(), len(s.s))
	if d.Len() >= len(s.s) {
		t.Errorf("TestCopyWithContext full copy")
	}

}
