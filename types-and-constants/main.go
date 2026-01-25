package main

import (
	"fmt"
	"strconv"
)

// bool
//
// int, int8, int16, int32, int64
// uint, uint8, uint16, uint32, uint64, uintptr

// byte = uint8
// rune = int32

// float32, float64
// complex64, complex128

func main() {
	var x int = 65
	f := string(x)
	fmt.Println(f)

	s := strconv.FormatInt(int64(x), 10)
	fmt.Println(s)

	const y int = 10
	fmt.Println(y)

	// untyped constant can be used as any type
	const z = 20
	takeInt32(z)
	takeInt64(z)
}

func takeInt32(x int32) {
	fmt.Println(x)
}

func takeInt64(x int64) {
	fmt.Println(x)
}
