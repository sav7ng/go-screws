package calendar

import (
	"fmt"
	"testing"
	"time"
)

func TestDmp(t *testing.T) {
	tool := &CalendarTool{}
	now := time.Now()
	fmt.Println(tool.FirstDayOfYear(now))
	fmt.Println(tool.LastDayOfYear(now))
	fmt.Println(tool.FirstDayOfMonth(now))
	fmt.Println(tool.LastDayOfMonth(now))
	fmt.Println(tool.FirstDayOfWeek(now))
	fmt.Println(tool.LastDayOfWeek(now))
	fmt.Println(tool.SomeDayBefore(now, 7))
	fmt.Println(tool.SomeDayAfter(now, 7))
}
