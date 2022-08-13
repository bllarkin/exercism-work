package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	timeLayout := "1/02/2006 15:04:05"
	initTime, _ := time.Parse(timeLayout, date)

	// fmt.Printf("\nTime Parsed\n%s\n", initTime.Local().String())

	return time.Date(initTime.Year(), initTime.Month(), initTime.Day(), initTime.Hour(), initTime.Minute(), initTime.Second(), initTime.Nanosecond(), initTime.Location())
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	timeLayout := "January 2, 2006 15:04:05"

	initTime, _ := time.Parse(timeLayout, date)

	return initTime.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	panic("Please implement the AnniversaryDate function")
}
