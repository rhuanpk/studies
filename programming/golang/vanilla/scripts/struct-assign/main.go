package main

import (
	"time"

	"module/structure"
)

// StructA struct A for tests.
type StructA struct {
	Field1 string
	Field2 int
	Field3 any
}

// StructB struct B for tests.
type StructB struct {
	Field1 any
	Field2 bool
	Field3 time.Time
	Field4 float32
}

func main() {
	objectB := &StructB{
		Field1: "hello",
		Field2: true,
		Field3: time.Now(),
		Field4: 42,
	}
	objectA := &StructA{}

	structure.Assign(objectA, objectB)
}
