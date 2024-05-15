package time

import (
	"errors"
	"time"
)

var ErrRange = errors.New("time range is error")

type Range struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

func (r Range) PeriodsOverlap(t Range) bool {
	return periodsOverlap(r, t)
}

func (r Range) IsEqual() bool {
	return r.StartTime.Equal(r.EndTime)
}

// Contains 判断一个时间是否在这个时间段内，包含临界时间
func (r Range) Contains(t time.Time) bool {
	return r.StartTime.Equal(t) || r.EndTime.Equal(t) || (r.StartTime.Before(t) && r.EndTime.After(t))
}

// PeriodsOverlap 判断两个时间段是否有重叠
func PeriodsOverlap(st Range, et Range) bool {
	return periodsOverlap(st, et)
}

func periodsOverlap(st Range, et Range) bool {
	return et.StartTime.Before(st.EndTime) && st.StartTime.Before(et.EndTime)
}

func NewRange(st time.Time, et time.Time) (Range, error) {
	if !st.Before(et) && !st.Equal(et) {
		return Range{}, ErrRange
	}
	return Range{StartTime: st, EndTime: et}, nil
}

func NewRangeFromString(st, et, layout string) (_ Range, err error) {
	return NewRangeFromStringInLocation(st, et, layout, time.Local)
}

func NewRangeFromStringInLocation(st, et, layout string, loc *time.Location) (_ Range, err error) {
	var t1, t2 time.Time
	t1, err = time.ParseInLocation(layout, st, loc)
	if err != nil {
		return
	}
	t2, err = time.ParseInLocation(layout, et, loc)
	if err != nil {
		return
	}
	return NewRange(t1, t2)
}
