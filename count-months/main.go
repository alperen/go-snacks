package main

import (
	"log"
	"math"
	"time"
)

const (
	hoursInDay   = 24
	daysInMonth  = 30
	hoursInMonth = daysInMonth * hoursInDay
)

func countMonths(t time.Time) float64 {
	return math.Floor(time.Since(t).Hours() / hoursInMonth)
}

func main() {
	date := time.Date(2019, 2, 2, 0, 0, 0, 0, time.UTC)

	log.Println(countMonths(date))
}
