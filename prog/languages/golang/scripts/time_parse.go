package main

import (
	"log"
	"time"
)

// go playground: https://go.dev/play/p/Izo9KHKbdhD

func main() {
	/*
		The pattern `2006-01-02 15:04:05` accepts:
			- yyyy-mm-dd hh:mm:ss
		The pattern `2006-01-02T15:04:05Z07:00` accepts:
			- yyyy-mm-ddThh:mm:ss{Z|{+|-}hh:mm}
		The pattern `2006-01-02 03:04:05 PM` accepts:
			- yyyy-mm-dd hh:mm:ss {AM|PM}
		The pattern `2006-01-02 03:04:05 PM 'Z07:00'` accepts:
			- yyyy-mm-dd hh:mm:ss {AM|PM} '{Z|{+|-}hh:mm}'
	*/
	layoutsXdatetimes := map[string]string{
		"2006-01-02 15:04:05":             "2023-08-30 20:57:35",
		"2006-01-02T15:04:05Z07:00":       "2023-08-30T20:57:35-03:00",
		"2006-01-02 03:04:05 PM":          "2023-08-30 08:57:35 AM",
		"2006-01-02 03:04:05 PM 'Z07:00'": "2023-08-30 08:57:35 PM '+05:00'",
	}
	for layout, datetime := range layoutsXdatetimes {
		println("layout:", layout, "->", "datetime:", datetime)
		final, err := time.Parse(layout, datetime)
		if err != nil {
			log.Fatalln(err)
		}
		println("-", final.String(), "\n")
	}
	println("OBS: Omitting the time zone will imply using the default (UTC).")
}
