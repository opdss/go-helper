package funcs

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRunTimer(t *testing.T) {
	n := 0
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	err := RunTimer(ctx, time.Second*3, func() error {
		n++
		return nil
	})
	if err != nil {
		t.Errorf("RunTimer error: %v", err)
		return
	}
	if n != 1 {
		t.Errorf("RunTimer error: fn not run")
	}
}
func TestRunTicker(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	n := 0
	err := RunTicker(ctx, time.Second*3, func() error {
		if n == 0 {
			time.Sleep(time.Second * 3)
		}
		n++
		return nil
	})

	if err != nil && !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("TestRunTicker error: %v", err)
		return
	}
	if n != 3 {
		t.Errorf("TestRunTicker error: fn run error")
	}
}

func TestRunTickerAfter(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	n := 0
	err := RunTickerAfter(ctx, time.Second*3, func() error {
		if n == 0 {
			time.Sleep(time.Second * 3)
		}
		n++
		return nil
	})

	if err != nil && !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("TestRunTicker error: %v", err)
		return
	}
	if n != 2 {
		t.Errorf("TestRunTicker error: fn run error")
	}

}
