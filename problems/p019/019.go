package p019

import (
	"github.com/washiil/euler-go/problems/registry"
)

/*
1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.

How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
*/

func isCenturialYear(year int) bool {
	return year%100 == 0
}

func isLeapYear(year int) bool {
	if isCenturialYear(year) {
		return year%400 == 0
	} else {
		return year%4 == 0
	}
}

func getDaysInMonth(year int, month int) int {
	days := month_lengths[month]
	if month == 2 && isLeapYear(year) {
		days += 1
	}
	return days
}

var month_lengths = map[int]int{
	1:  31, // January
	2:  28, // Febuary
	3:  31, // ...
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31, // ...
	11: 30, // November
	12: 31, // December
}

func Solve019() int {
	sundays := 0 // Output variable

	month := 1       // January
	year := 1900     // 1900 (duh)
	day_of_week := 1 // Monday

	// Loop until we reach end year
	for year < 2001 {
		// Note how we start at 1900 to establish correct dates but dont count until year 1901
		if year >= 1901 && day_of_week == 0 {
			sundays++
		}

		days_in_month := getDaysInMonth(year, month)
		// Nice modulous math to simplify operations and loops
		day_of_week = (day_of_week + days_in_month) % 7

		month++
		if month > 12 {
			month = 1
			year++
		}
	}

	return sundays
}

func init() {
	registry.Register("19", Solve019)
}
