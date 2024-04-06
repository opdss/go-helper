package funcs

import (
	"context"
	"time"
)

// RunTimer 定时器执行一次
func RunTimer(ctx context.Context, t time.Duration, fn func() error) error {
	tt := time.NewTimer(t)
	defer tt.Stop()
	select {
	case <-tt.C:
		return fn()
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RunTicker 定时器间隔执行
func RunTicker(ctx context.Context, t time.Duration, fn func() error) error {
	tk := time.NewTicker(t)
	defer tk.Stop()
	for {
		select {
		case <-tk.C:
			if err := fn(); err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// RunTickerAfter 定时器间隔执行,在每次执行完后，在开始定时下一次执行
func RunTickerAfter(ctx context.Context, t time.Duration, fn func() error) error {
	tt := time.NewTimer(t)
	defer tt.Stop()
	for {
		select {
		case <-tt.C:
			if err := fn(); err != nil {
				return err
			}
			tt.Reset(t)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
