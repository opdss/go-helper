package time

import (
	"errors"
	"testing"
	"time"
)

func TestPeriodsOverlap(t *testing.T) {
	t1, _ := time.Parse(time.DateTime, "2022-02-22 12:12:12")
	t2, _ := time.Parse(time.DateTime, "2022-02-22 13:12:12")
	t3, _ := time.Parse(time.DateTime, "2022-02-22 14:12:12")
	t4, _ := time.Parse(time.DateTime, "2022-02-22 15:12:12")

	r1 := Range{StartTime: t1, EndTime: t2}
	r2 := Range{StartTime: t2, EndTime: t3}
	r3 := Range{StartTime: t3, EndTime: t4}
	r4 := Range{StartTime: t1, EndTime: t3}
	r5 := Range{StartTime: t2, EndTime: t4}
	r6 := Range{StartTime: t1, EndTime: t4}
	r7 := Range{StartTime: t2, EndTime: t3}
	arr := map[[2]Range]bool{
		[2]Range{r1, r3}: false,
		[2]Range{r1, r2}: false,
		[2]Range{r4, r5}: true,
		[2]Range{r6, r7}: true,
		[2]Range{r7, r6}: true,
	}

	for k, v := range arr {
		if res := PeriodsOverlap(k[0], k[1]); res != v {
			t.Errorf("PeriodsOverlap(%+v, %+v) == %v, Want %v", k[0], k[1], res, v)
		}
	}
}

func TestNewRangeFromString(t *testing.T) {
	_, err := NewRangeFromString("2022-02-22 12:12:12", "2022-02-22 13:12:12", time.DateTime)
	if err != nil {
		t.Errorf("TestNewRangeFromString error:%v", err)
		return
	}

	_, err = NewRangeFromString("2022-02-22 13:12:12", "2022-02-22 12:12:12", time.DateTime)
	if !errors.Is(err, ErrRange) {
		t.Errorf("TestNewRangeFromString error:%v", err)
	}
}
