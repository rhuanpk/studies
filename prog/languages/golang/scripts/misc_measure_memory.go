package main

import (
	"math"
	"math/rand"
	"runtime"
	"strconv"
)

// golang memory
//
// stack:
//  - variáveis locais de função (que não retornam/escapam)
//  - parâmetros de função
//  - valores de retorno primitivos
//
// heap:
//  - variáveis que retornam/escapam de função
//  - objetos criados com make() ou new()
//  - slices, maps, channels, structs
//  - ponteiros que retornam/saem de função
//
// https://claude.ai/share/04c7f592-4e43-4d1e-8806-23d524d1f2f5
//
// see also: https://github.com/go-hl/memory

func toMiB(mem uint64) uint64 {
	return mem / uint64(math.Pow(1024, 2))
}

func mem() {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	println("alloc:", toMiB(stats.Alloc), "MiB")      // current memory in heap
	println("sys:", toMiB(stats.Sys), "MiB")          // current reserved memory to go runtime in os
	println("total:", toMiB(stats.TotalAlloc), "MiB") // total memory in heap during all program exection
	println("gc:", stats.NumGC, "c")                  // total garbage collector cycles
}

func sum() {
	var sum int
	sumMap := make(map[string]int)

	for range 999999 {
		value := rand.Int()
		key := strconv.Itoa(value)
		sumMap[key] = value
	}

	for _, value := range sumMap {
		sum += value
	}

	println("sum:", sum)
}

func main() {
	sum()
	mem()
	// sum() // increase total alloc
	// sum() // increase total alloc
	// mem() // see stats again
}
