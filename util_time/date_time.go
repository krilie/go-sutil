package util_time

import (
	"math"
	"time"
)

const TimeFormatNormal = "2006-01-02 15:04:05"
const TimeFormatRFC3339 = time.RFC3339
const TimeFormatDate = "2006-01-02"
const TimeFormatTime = "15:04:05"

// 本月开始时间
func GetBeijingMonthStartTime(time time.Time) time.Time {
	time = time.In(BeijingZone).AddDate(0, 0, -time.Day()+1)
	return GetBeijingZeroTime(time)
}

// 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetBeijingLastDateOfMonth(d time.Time) time.Time {
	return GetBeijingMonthStartTime(d).In(BeijingZone).AddDate(0, 1, 0).Add(-time.Second)
}

// 获取零点时间
func GetBeijingZeroTime(d time.Time) time.Time {
	d = d.In(BeijingZone)
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

func GetTodayStartTime() time.Time {
	return GetZeroTime(time.Now())
}
func GetTodayEndTime() time.Time {
	return GetEndTime(time.Now())
}
func GetYesterdayStartTime() time.Time {
	return GetZeroTime(time.Now().AddDate(0, 0, -1))
}
func GetYesterdayEndTime() time.Time {
	return GetEndTime(time.Now().AddDate(0, 0, -1))
}

func GetFirstDateOfPreMonth(d time.Time) time.Time {
	d = GetFirstDateOfMonth(GetFirstDateOfMonth(d).AddDate(0, 0, -1))
	return GetZeroTime(d)
}
func GetLastDateOfPreMonth(d time.Time) time.Time {
	return GetLastDateOfMonth(GetFirstDateOfMonth(d).AddDate(0, 0, -1))
}

func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetEndTime(GetFirstDateOfMonth(d).AddDate(0, 1, -1))
}

func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}
func GetBeginTime(d time.Time) time.Time {
	return GetZeroTime(d)
}
func GetEndTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 0, d.Location())
}

// GetWeekStart 取到本周的开始时间
func GetWeekStart(d time.Time) time.Time {
	offset := int(time.Monday - d.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStart := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	return weekStart
}

// GetWeekEnd 取到本周的结束时间 0 1 2 3 4 5 6
func GetWeekEnd(d time.Time) time.Time {
	offset := int(time.Saturday - d.Weekday())
	weekEnd := time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 0, time.Local).AddDate(0, 0, offset+1)
	return weekEnd
}

// DaysOnTwoTime 计算两个时间相差的天数
func DaysOnTwoTime(begin, end time.Time) int {
	// 计算两个时间点的相差天数
	dt1 := time.Date(begin.Year(), begin.Month(), begin.Day(), 0, 0, 0, 0, time.Local)
	dt2 := time.Date(end.Year(), end.Month(), end.Day(), 0, 0, 0, 0, time.Local)
	return int(math.Ceil(dt2.Sub(dt1).Hours() / 24))
}

// ParseBeiJingTimeStr 解析时间,解析不成功 返回nil
func ParseBeiJingTimeStr(str string, format string) *time.Time {
	if format == "" {
		format = TimeFormatNormal
	}
	parse, err := time.Parse(format, str)
	if err != nil {
		return nil
	} else {
		return &parse
	}
}

func TimeToStr(t time.Time, format string) string { return t.Format(format) }
func PtrTime(t time.Time) *time.Time              { return &t }
func ParseTimeOrNil(tstr string, format string) *time.Time {
	parse, err := time.Parse(format, tstr)
	if err != nil {
		return nil
	}
	return &parse
}
func ParseTimeOrNow(tstr string, format string) time.Time {
	parse, err := time.Parse(format, tstr)
	if err != nil {
		return time.Now()
	}
	return parse
}
