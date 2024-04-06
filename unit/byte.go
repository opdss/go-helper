package unit

import (
	"fmt"
	"strconv"
	"time"
)

// ByteFormat 将byte转成带单位的字面量
func ByteFormat(n int64, precision uint8) string {
	fs := [6]string{"B", "KB", "MB", "GB", "TB", "PB"}
	r := float64(n)
	u := fs[0]
	for i := range fs {
		if r < 1024 || i == 5 {
			break
		}
		u = fs[i+1]
		r = r / 1024
	}
	return fmt.Sprintf("%."+strconv.Itoa(int(precision))+"f%s", r, u)
}

// NetworkSpeed 将byte单位带宽，转成带单位的字面量
func NetworkSpeed(size int64, useTime time.Duration, precision uint8) string {
	fs := [4]string{"bps", "Kbps", "Mbps", "Gbps"}
	r := float64(size) * 8 / useTime.Seconds()
	u := fs[0]
	for i := range fs {
		if r < 1e3 || i == 3 {
			break
		}
		u = fs[i+1]
		r = r / 1e3
	}
	return fmt.Sprintf("%."+strconv.Itoa(int(precision))+"f%s", r, u)
}
