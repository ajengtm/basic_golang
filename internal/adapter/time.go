package adapter

import (
	"time"
)

var (
	DefaultTimeZoneLocation string = "Asia/Jakarta"
	YYYYMMDDFormat          string = "2006-01-02"
	YYYYMMDDHHMMSSFormat    string = "2006-01-02 15:04:05"
	DDMMYYYYHHMMSSFormat    string = "02-01-2006 15:04:05"
)

func GetCurrentTimestamp() (timestamp string) {
	loc, _ := time.LoadLocation(DefaultTimeZoneLocation)
	now := time.Now().In(loc)
	return now.Format(YYYYMMDDHHMMSSFormat)
}

func GetCurrentTimestampTZ() (timestamp string) {
	loc, _ := time.LoadLocation(DefaultTimeZoneLocation)
	now := time.Now().In(loc)
	return now.Format("2006-01-02T15:04:05")
}

func GetYesterday() (timestamp string) {
	loc, _ := time.LoadLocation(DefaultTimeZoneLocation)
	return time.Now().In(loc).AddDate(0, 0, -1).Format("2006-01-02")
}

func GetNowWithFormat(format string) (timestamp string) {
	loc, _ := time.LoadLocation(DefaultTimeZoneLocation)
	now := time.Now().In(loc)
	return now.Format(format)
}

func GetCurrentDate() (timestamp string) {
	loc, _ := time.LoadLocation(DefaultTimeZoneLocation)
	now := time.Now().In(loc)
	return now.Format("2006-01-02")
}

func GeneralFormatTime(in string, layout string) string {
	loc, _ := time.LoadLocation(DefaultTimeZoneLocation)
	t, err := time.ParseInLocation(time.RFC3339, in, loc)
	if err != nil {
		return in
	}
	return t.Format(layout)
}

func FormatTimestamp(in string) string {
	loc, _ := time.LoadLocation(DefaultTimeZoneLocation)
	t, err := time.ParseInLocation(time.RFC3339, in, loc)
	if err != nil {
		return in
	}
	return t.Format(YYYYMMDDHHMMSSFormat)
}

func FormatDate(in string) string {
	loc, _ := time.LoadLocation(DefaultTimeZoneLocation)
	t, err := time.ParseInLocation(time.RFC3339, in, loc)
	if err != nil {
		return in
	}
	return t.Format("2006-01-02")
}

// FirstDayOfISOWeek return first day timestamp for given timestamp
func FirstDayOfISOWeek(year int, week int) time.Time {
	loc, _ := time.LoadLocation(DefaultTimeZoneLocation)
	date := time.Date(year, 0, 0, 0, 0, 0, 0, loc)
	isoYear, isoWeek := date.ISOWeek()
	for date.Weekday() != time.Monday { // iterate back to Monday
		date = date.AddDate(0, 0, -1)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoYear < year { // iterate forward to the first day of the first week
		date = date.AddDate(0, 0, 1)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoWeek < week { // iterate forward to the first day of the given week
		date = date.AddDate(0, 0, 1)
		isoYear, isoWeek = date.ISOWeek()
	}
	return date
}

// IsValidFormatTimestamp check whether inputted timestamp is match the desired format or not
func IsValidFormatTimestamp(inputTimestamp string, desiredFormat string) bool {
	_, err := time.Parse(desiredFormat, inputTimestamp)
	if err != nil {
		return false
	}
	return true
}

func GetCurrentTimestampFromString(timestamp string) time.Time {
	loc, _ := time.LoadLocation(DefaultTimeZoneLocation)
	t, _ := time.ParseInLocation(YYYYMMDDHHMMSSFormat, timestamp, loc)
	return t
}

func BeginningOfDay(t time.Time) time.Time {
	loc, _ := time.LoadLocation(DefaultTimeZoneLocation)
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location()).In(loc)
}

func EndOfDay(t time.Time) time.Time {
	loc, _ := time.LoadLocation(DefaultTimeZoneLocation)
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, t.Location()).In(loc)
}
