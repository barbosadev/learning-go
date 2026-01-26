package main

import "fmt"

func main() {
	if 1 < 2 {
		fmt.Println("1 is less than 2")
	}

	// I can declare a variable in the if statement
	// I only have access to x inside the if block
	if x := 10; x < 10 {
		fmt.Println(x, "is less than or equal to 10")
	} else if x > 11 {
		fmt.Println(x, "is greater than 11")
	} else {
		fmt.Println(x, "is between 10 and 11")
	}
}
