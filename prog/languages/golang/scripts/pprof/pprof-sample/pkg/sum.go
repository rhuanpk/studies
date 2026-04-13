package pkg

import (
	"math/rand"
	"strconv"
)

func CreazySum(start, end int) {
	var sum int
	sumMap := make(map[string]int)

	for index := start; index < end; index++ {
		sumMap[strconv.Itoa(rand.Int())] = index
	}

	for _, value := range sumMap {
		sum += value
	}

	println("total:", sum)
}
