package funcs

import (
	"context"
	"io"
)

type readerFunc func(p []byte) (n int, err error)

func (rf readerFunc) Read(p []byte) (n int, err error) { return rf(p) }

// CopyWithContext 可以context取消的io.copy
func CopyWithContext(ctx context.Context, dst io.Writer, src io.Reader) (n int64, err error) {
	n, err = io.Copy(dst, readerFunc(func(p []byte) (int, error) {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		default:
			return src.Read(p)
		}
	}))
	return
}
