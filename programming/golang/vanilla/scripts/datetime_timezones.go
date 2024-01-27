package main

import (
	"fmt"
	"time"
)

func customLocation() {
	println("---------- Custom Location ----------")

	println("\n***** Before *****")

	println("\n> zoneName, _ := time.Now().In(time.Now().Location()).Zone()")
	zoneName, _ := time.Now().In(time.Now().Location()).Zone()
	println("\n> fmt.Println(zoneName)")
	fmt.Println(zoneName)

	println("\n> time.Now().Location()")
	fmt.Println(time.Now().Location())

	println("\n> time.Now()")
	fmt.Println(time.Now())

	println("\n***** After *****")

	println("\n> location, _ := time.LoadLocation(\"Asia/Tokyo\")")
	location, _ := time.LoadLocation("Asia/Tokyo")

	println("\n> zoneName, _ = time.Now().In(location).Zone()")
	zoneName, _ = time.Now().In(location).Zone()
	println("\n> fmt.Println(zoneName)")
	fmt.Println(zoneName)

	println("\n> time.Now().In(location).Location()")
	fmt.Println(time.Now().In(location).Location())

	println("\n> time.Now().In(location)")
	fmt.Println(time.Now().In(location))

	println("\n***** Like Back *****")

	println("\n> zoneName, _ = time.Now().In(time.Now().Location()).Zone()")
	zoneName, _ = time.Now().In(time.Now().Location()).Zone()
	println("\n> fmt.Println(zoneName)")
	fmt.Println(zoneName)

	println("\n> time.Now().Location()")
	fmt.Println(time.Now().Location())

	println("\n> time.Now()")
	fmt.Println(time.Now())
}

func standardLocation() {
	println("---------- Standard Location ----------")

	println("\n***** Before *****")

	println("\n> zoneName, _ := time.Now().In(time.Now().Location()).Zone()")
	zoneName, _ := time.Now().In(time.Now().Location()).Zone()
	println("\n> fmt.Println(zoneName)")
	fmt.Println(zoneName)

	println("\n> time.Now().Location()")
	fmt.Println(time.Now().Location())

	println("\n> time.Now()")
	fmt.Println(time.Now())

	println("\n***** After *****")

	println("\n> time.Local, _ = time.LoadLocation(\"America/Toronto\")")
	time.Local, _ = time.LoadLocation("America/Toronto")

	println("\n> zoneName, _ := time.Now().In(time.Now().Location()).Zone()")
	zoneName, _ = time.Now().In(time.Now().Location()).Zone()
	println("\n> fmt.Println(zoneName)")
	fmt.Println(zoneName)

	println("\n> time.Now().Location()")
	fmt.Println(time.Now().Location())

	println("\n> time.Now()")
	fmt.Println(time.Now())
}

func fixedZone() {
	println("---------- Fixed Zone ----------")

	println("\n***** Before *****")

	println("\n> zoneName, _ := time.Now().In(time.Now().Location()).Zone()")
	zoneName, _ := time.Now().In(time.Now().Location()).Zone()
	println("\n> fmt.Println(zoneName)")
	fmt.Println(zoneName)

	println("\n> time.Now().Location()")
	fmt.Println(time.Now().Location())

	println("\n> time.Now()")
	fmt.Println(time.Now())

	println("\n***** After *****")

	println("\n> zoneName = \"Europe/Istanbul\"")
	zoneName = "Europe/Istanbul"

	println("\n> location, _ := time.LoadLocation(zoneName)")
	location, _ := time.LoadLocation(zoneName)

	println("\n> _, offset := time.Now().In(location).Zone()")
	_, offset := time.Now().In(location).Zone()

	println("\n> fmt.Println(time.Now().In(time.FixedZone(zoneName, offset)))")
	fmt.Println(time.Now().In(time.FixedZone(zoneName, offset)))

	println("\n***** Like Back *****")

	println("\n> zoneName, _ := time.Now().In(time.Now().Location()).Zone()")
	zoneName, _ = time.Now().In(time.Now().Location()).Zone()
	println("\n> fmt.Println(zoneName)")
	fmt.Println(zoneName)

	println("\n> time.Now().Location()")
	fmt.Println(time.Now().Location())

	println("\n> time.Now()")
	fmt.Println(time.Now())
}

// go playground: https://go.dev/play/p/GLOoxgW9gRl
func main() {
	customLocation()
	println()
	standardLocation()
	println()
	fixedZone()
}
