package utils

import "time"

func CreateDateFromMonthAndYear(month, year int) string {
	// Creating a date with the given month and year
	date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

	// Formatting the date to YYYY-MM-DD format
	return date.Format("2006-01-02")
}
