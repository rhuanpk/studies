package main

// Playground: https://go.dev/play/p/zg7UjwIg3uF
func main() {
	// Only map declaration.
	var mapa1 map[string]any
	if len(mapa1) == 0 {
		println("empty")
	} else {
		println("filled")
	}

	// Post inicialization.
	var mapa2 map[string]any
	mapa2 = make(map[string]any)
	mapa2["hello"] = "world"
	if len(mapa2) == 0 {
		println("empty")
	} else {
		println("filled")
	}

	// Inicialization on declaration.
	mapa3 := map[string]any{"hello": "world"}
	if len(mapa3) == 0 {
		println("empty")
	} else {
		println("filled")
	}
}
