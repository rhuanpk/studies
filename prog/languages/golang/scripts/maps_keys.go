package main

import (
	"fmt"

	"golang.org/x/exp/maps"
)

// go playground: https://go.dev/play/p/3Te4l1yRoou
func main() {
	mapa := map[any]any{
		"hello": "world",
		42:      "every",
	}
	fmt.Printf("map keys: %#v\n", maps.Keys(mapa))
}
