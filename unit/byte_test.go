package unit

import (
	"testing"
	"time"
)

var num = int64(1024)
var num2 = int64(2000000)

func TestByteFormat(t *testing.T) {
	s := ByteFormat(num, 3)
	v := "1.000KB"
	if s != v {
		t.Errorf("ByteFormat(%d) = %s, want %s", num, s, v)
	}
}

func TestNetworkSpeed(t *testing.T) {
	s := NetworkSpeed(num2, time.Second*2, 3)
	v := "8.000Mbps"
	if s != v {
		t.Errorf("ByteFormat(%d) = %s, want %s", num, s, v)
	}
}
