package time

import (
	"math"
	"time"
)

type TimeTool struct{}

var Zone = time.FixedZone("GMT", 8)

const Layout = "2006-01-02 15:04:05"

func (s *TimeTool) ToTime(value string) *time.Time {
	t, err := time.ParseInLocation(Layout, value, Zone)
	if err != nil {
		return nil
	}
	return &t
}

func (s *TimeTool) FromTime(value time.Time) string {
	return value.In(Zone).Format(Layout)
}

// 计算时间差（天数）
func (s *TimeTool) SubDays(t1, t2 time.Time) (day int) {
	return int(math.Abs(t1.Sub(t2).Hours() / 24))
}

// 计算时间差（天数） 四舍五入
func (s *TimeTool) SubDaysRound(t1, t2 time.Time) (day int) {
	return int(math.Abs(math.Round(t1.Sub(t2).Hours() / 24)))
}
