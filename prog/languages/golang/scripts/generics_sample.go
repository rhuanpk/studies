package main

import "fmt"

// go playground: https://go.dev/play/p/WMKQF1vlLsp
// pastebin: https://pastebin.com/4Aipbqp3

type mainStruct[Type any] struct {
	firstField  string
	secondField bool
	slaveStruct slaveStruct
	Struct      Type
}

type slaveStruct struct {
	firstField  string
	secondField bool
}

type otherStruct struct {
	firstField  string
	secondField bool
}

func makeStruct[Type any](source *Type) *mainStruct[Type] {
	return &mainStruct[Type]{
		firstField:  "",
		secondField: false,
		slaveStruct: slaveStruct{},
		Struct:      *source,
	}
}

func main() {
	fullStruct := makeStruct(&otherStruct{})
	fmt.Printf("%+v\n", *fullStruct)
}
