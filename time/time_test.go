package time

import (
	"testing"
	"time"
)

func TestSubDays(t *testing.T) {
	tool := &TimeTool{}
	timeFormat := "2006-01-02 15:04:05" // 默认转换日期格式
	beginTime, err := time.Parse(timeFormat, "2023-11-08 13:47:59")
	if err != nil {
		panic(err)
	}
	subDays := tool.SubDays(time.Now(), beginTime)
	println(subDays)
}
