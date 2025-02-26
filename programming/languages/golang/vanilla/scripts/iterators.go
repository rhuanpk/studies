package main

import (
	"fmt"
	"iter"
	"slices"
)

// see also: slices.Sorted(), slices.SortedFunc(), slices.SortedStableFunc(), slices.Chunk(), maps.*
// go playground: https://go.dev/play/p/L-Mt0nmZhrk
func main() {
	array := []any{0, 1, "hello", "world", "xpto"}

	println("-----> for range value")
	for _, value := range array {
		fmt.Printf("value: %#v\n", value)
	}
	println("-----> for range slices.Values()")
	for value := range slices.Values(array) {
		fmt.Printf("value: %#v\n", value)
	}
	println("-----> slices.Values() seq")
	slices.Values(array)(func(value any) bool {
		fmt.Printf("value: %#v\n", value)
		return true
	})

	println("\n-----> for range index, value")
	for index, value := range array {
		fmt.Printf("key: %d value: %#v\n", index, value)
	}
	println("-----> for range slices.All()")
	for index, value := range slices.All(array) {
		fmt.Printf("key: %d value: %#v\n", index, value)
	}
	println("-----> slices.All() seq2")
	slices.All(array)(func(index int, value any) bool {
		fmt.Printf("key: %d value: %#v\n", index, value)
		return true
	})

	println("\n-----> for range -index, -value")
	for index := len(array) - 1; index >= 0; index-- {
		fmt.Printf("key: %d value: %#v\n", index, array[index])
	}
	println("-----> for range slices.Backward()")
	for index, value := range slices.Backward(array) {
		fmt.Printf("key: %d value: %#v\n", index, value)
	}
	println("-----> slices.Backward() seq2")
	slices.Backward(array)(func(index int, value any) bool {
		fmt.Printf("key: %d value: %#v\n", index, value)
		return true
	})

	println("\n-----> for range value append")
	for _, value := range append([]any{-1.2}, array...) {
		fmt.Printf("value: %#v\n", value)
	}
	println("-----> for range slices.AppendSeq()")
	for _, value := range slices.AppendSeq([]any{-1.2}, slices.Values(array)) {
		fmt.Printf("value: %#v\n", value)
	}

	println("\n-----> iterator1")
	iterator := func(number int) iter.Seq[int] {
		return func(yield func(int) bool) {
			for index := range number {
				if !yield(index) {
					return
				}
			}
		}
	}

	sequence1 := iterator(5)
	sequence1(func(value int) bool {
		fmt.Printf("value: %#v\n", value)
		// `return false' breaks the loop
		return true
	})

	println("-----> iterator2")
	iterator2 := func(kv map[string]any) iter.Seq2[string, any] {
		return func(yield func(string, any) bool) {
			for key, value := range kv {
				if !yield(key, value) {
					return
				}
			}
		}
	}
	sequence2 := iterator2(map[string]any{"foo": "bar", "hello": "world"})
	sequence2(func(key string, value any) bool {
		fmt.Printf("key: %s, value: %#v\n", key, value)
		// `return false' breaks the loop
		return true
	})
}
