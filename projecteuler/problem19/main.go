package main

import (
	"fmt"
	"time"
)

const (
	centuryStart = 1901
	centuryEnd   = 2000
)

var totalSundays int

func main() {

	for year := centuryStart; year <= centuryEnd; year++ {

		for month := 1; month <= 12; month++ {

			then := time.Date(
				year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
			if then.Weekday() == time.Sunday {
				//fmt.Println(then.Year(), then.Weekday(), then.Month(), then.Day())
				totalSundays++
			}
		}
	}
	fmt.Println("The total number of Sundays that fell on the first of the month during the 20th century was: ", totalSundays)

}

/*

	https://projecteuler.net/problem=19

	You are given the following information, but you may prefer to do some research for yourself.

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