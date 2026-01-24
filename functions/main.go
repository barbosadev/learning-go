package main

import "fmt"

func main() {
	helloWorld()
}

func helloWorld() {
	fmt.Println("Hello, World!")

	fmt.Println(sum(8, 5))
	fmt.Println(swap(8, 5))

	res, rem := split(17, 3)
	fmt.Println(res, rem)

	fmt.Println(multiply(4)(5))        // Outputs 20
	fmt.Println(sumAll(1, 2, 3, 4, 5)) // Outputs 15
}

func sum(a int, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

func split(a, b int) (res int, rem int) {
	res = a / b
	rem = a % b

	return res, rem
	// I can use a naked return here,
	// but it's not recommended. ex.:
	// return
}

// Higher-order function example
func multiply(a int) func(int) int {
	return func(b int) int {
		return a * b
	}
}

// Variadic function example
func sumAll(numbers ...int) int {
	var total int
	for _, n := range numbers {
		total += n
	}

	return total
}
