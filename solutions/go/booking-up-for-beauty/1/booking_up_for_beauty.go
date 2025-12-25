package booking

import (
	"fmt"
	"time"
)

/*
	Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"

	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
	DateTime   = "2006-01-02 15:04:05"
	DateOnly   = "2006-01-02"
	TimeOnly   = "15:04:05"

	Year: "2006" "06"
	Month: "Jan" "January" "01" "1"
	Day of the week: "Mon" "Monday"
	Day of the month: "2" "_2" "02"
	Day of the year: "__2" "002"
	Hour: "15" "3" "03" (PM or AM)
	Minute: "4" "04"
	Second: "5" "05"
	AM/PM mark: "PM"
*/

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		return time.Time{}
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		return false
	}
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		return false
	}
	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
// Description("7/25/2019 13:45:00")
// => "You have an appointment on Thursday, July 25, 2019, at 13:45."
func Description(date string) string {
	t, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("You have an appointment on %s, %s, at %s.", t.Weekday(), t.Format("January 2, 2006"), t.Format("15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
