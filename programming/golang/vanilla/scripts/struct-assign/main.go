package main

import (
	"time"

	"module/structure"
)

type structA struct {
	Field1 string
	Field2 int
	Field3 any
	field4 float64
	Field5 interface{}
}

type structB struct {
	Field1 any
	Field2 bool
	Field3 time.Time
	Field4 float32
}

func main() {
	objectA := &structA{}
	objectB := &structB{
		Field1: "hello",
		Field2: true,
		Field3: time.Now(),
		Field4: 42,
	}

	structure.Assign(objectB, objectA)
}
