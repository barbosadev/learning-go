package main

import "fmt"

func main() {
	// for loops
	var x int
	for i := 0; i < 10; i++ {
		x++
	}
	fmt.Println(x)

	var i int
	for i < 10 {
		x++
		i++
	}
	fmt.Println(x)

	i = 0
	for {
		if i >= 10 {
			break
			// I can use break to exit the loop
		}
		x++
		i++
	}
	fmt.Println(x)

	// range loops
	arrx := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for j, elem := range arrx {
		fmt.Println(j, elem)
	}

	// blank identifier
	arry := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, elem := range arry {
		fmt.Println(elem)
	}

	for range 10 {
		fmt.Println("Hello!")
	}

	for k := range 10 {
		fmt.Println(k)
	}

	arrz := [3]int{1, 2, 3}
	for i, elem := range arrz {
		fmt.Println(&i, &elem)
	}
}
