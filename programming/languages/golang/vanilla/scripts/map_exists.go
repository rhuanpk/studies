package main

import "fmt"

// go playground: https://go.dev/play/p/eBTORWcJ17X
func main() {

	mapa := map[string]string{
		"primeiro": "zero",
		"segundo":  "",
	}

	var (
		primeiro, existsPrimeiro = mapa["primeiro"]
		segundo, existsSegundo   = mapa["segundo"]
		terceiro, existsTerceiro = mapa["terceiro"]
	)

	fmt.Printf(
		"primeiro, existsPrimeiro: %#v, %#v\nsegundo, existsSegundo: %#v, %#v\nterceiro, existsTerceiro: %#v, %#v\n",
		primeiro, existsPrimeiro, segundo, existsSegundo, terceiro, existsTerceiro,
	)

}
