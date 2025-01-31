package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time")

	presentTime := time.Now()
	fmt.Println("Present Time - ", presentTime)

	// The values present in the Format function are standard values,
	// to denote different parts of a date
	// the below format is as follows -
	// month-date-year hour:minutes:seconds day_of_week
	fmt.Println(presentTime.Format("01-02-2006 15-04-05 Monday"))

	// Specify a Date
	createdDate := time.Date(2024, time.June, 02, 22, 10, 04, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("Jan-02 2006 Monday"))
}
