package time

import (
	_ "embed"

	"github.com/ZenLiuCN/engine"
	"time"
)

var (
	//go:embed golang_time.d.ts
	TimeDefine   []byte
	TimeDeclared = map[string]any{
		"parseInLocation":        time.ParseInLocation,
		"newTimer":               time.NewTimer,
		"Hour":                   time.Hour,
		"Microsecond":            time.Microsecond,
		"Nanosecond":             time.Nanosecond,
		"Millisecond":            time.Millisecond,
		"Second":                 time.Second,
		"Minute":                 time.Minute,
		"after":                  time.After,
		"unix":                   time.Unix,
		"unixMilli":              time.UnixMilli,
		"unixMicro":              time.UnixMicro,
		"Monday":                 time.Monday,
		"Sunday":                 time.Sunday,
		"Tuesday":                time.Tuesday,
		"Wednesday":              time.Wednesday,
		"Thursday":               time.Thursday,
		"Friday":                 time.Friday,
		"Saturday":               time.Saturday,
		"loadLocation":           time.LoadLocation,
		"afterFunc":              time.AfterFunc,
		"tick":                   time.Tick,
		"February":               time.February,
		"May":                    time.May,
		"July":                   time.July,
		"August":                 time.August,
		"September":              time.September,
		"January":                time.January,
		"March":                  time.March,
		"April":                  time.April,
		"June":                   time.June,
		"October":                time.October,
		"November":               time.November,
		"December":               time.December,
		"since":                  time.Since,
		"now":                    time.Now,
		"parse":                  time.Parse,
		"sleep":                  time.Sleep,
		"until":                  time.Until,
		"newTicker":              time.NewTicker,
		"date":                   time.Date,
		"fixedZone":              time.FixedZone,
		"loadLocationFromTzData": time.LoadLocationFromTZData,
		"parseDuration":          time.ParseDuration,
	}
)

func init() {
	engine.RegisterModule(TimeModule{})
}

type TimeModule struct{}

func (S TimeModule) Identity() string {
	return "golang/time"
}
func (S TimeModule) TypeDefine() []byte {
	return TimeDefine
}
func (S TimeModule) Exports() map[string]any {
	return TimeDeclared
}
