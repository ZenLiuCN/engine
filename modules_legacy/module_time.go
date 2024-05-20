package legacy

import (
	_ "embed"
	"time"
)

var (
	//go:embed module_time.d.ts
	timeDefine []byte
	timeMap    = map[string]any{

		"January":   time.January,
		"February":  time.February,
		"March":     time.March,
		"April":     time.April,
		"May":       time.May,
		"June":      time.June,
		"July":      time.July,
		"August":    time.August,
		"September": time.September,
		"October":   time.October,
		"November":  time.November,
		"December":  time.December,
		"monthNumber": func(m time.Month) int {
			return int(m)
		},
		"numberMonth": func(m int) time.Month {
			return time.Month(m)
		},

		"Sunday":    time.Sunday,
		"Monday":    time.Monday,
		"Tuesday":   time.Tuesday,
		"Wednesday": time.Wednesday,
		"Thursday":  time.Thursday,
		"Friday":    time.Friday,
		"Saturday":  time.Saturday,

		"weekdayNumber": func(m time.Weekday) int {
			return int(m)
		},
		"numberWeekday": func(m int) time.Weekday {
			return time.Weekday(m)
		},
		"toDuration": func(n int, u time.Duration) time.Duration {
			return time.Duration(n) * u
		},
		"fromDuration": func(n, u time.Duration) int {
			return int(n / u)
		},
		"since":    time.Since,
		"duration": time.ParseDuration,
		"now":      time.Now,
		"parse":    time.Parse,
		"until":    time.Until,
		"times": func(d time.Duration, n int64) time.Duration {
			return d * time.Duration(n)
		},
		"fixedZone": time.FixedZone,
		"UTC":       time.UTC,
		"loadLocation": func(name string) *time.Location {
			l, _ := time.LoadLocation(name)
			return l
		},
		"date":        time.Date,
		"Nanosecond":  time.Nanosecond,
		"Microsecond": time.Microsecond,
		"Millisecond": time.Millisecond,
		"Second":      time.Second,
		"Minute":      time.Minute,
		"Hour":        time.Hour,
		"unix":        time.Unix,
		"unixMilli":   time.UnixMilli,
		"unixMicro":   time.UnixMicro,
	}
)

type TimeModule struct {
}

func (t TimeModule) Identity() string {
	return "go/time"
}

func (t TimeModule) TypeDefine() []byte {
	return timeDefine
}

func (t TimeModule) Exports() map[string]any {

	return timeMap
}
