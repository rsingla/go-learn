package main

import (
	"flag"
	"log"
	"time"
)

var expectedFormat = "2006-01-02"

// parseTime validates and parses a given date string.
func parseTime(target string) time.Time {
	time, err := time.Parse(expectedFormat, target)

	if err != nil {
		panic("Invalid time")
	}
	return time
}

// calcSleeps returns the number of sleeps until the target.
func calcSleeps(target time.Time) float64 {
	duration := time.Until(target)
	return float64(duration.Hours() / 24)
}

func main() {
	bday := flag.String("bday", "2023-12-30", "Your next bday in YYYY-MM-DD format")
	flag.Parse()
	target := parseTime(*bday)
	log.Printf("You have %d sleeps until your birthday. Hurray!",
		int(calcSleeps(target)))
}
