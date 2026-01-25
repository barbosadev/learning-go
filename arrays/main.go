package main

import "fmt"

func main() {
	arrx := [3]int{}
	fmt.Println(arrx)

	arry := [3]int{1, 2}
	fmt.Println(arry)

	arrz := [5]int{3: 10, 1: 5}
	fmt.Println(arrz)

	stringArray := [3]string{"hello", "world"}
	fmt.Println(stringArray)

	// arrays have fixed size
	// I can use slices for dynamic size
}
