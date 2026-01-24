package main

import "fmt"

var age int = 26

// var age = 26

func main() {
	var name, lastName string
	fmt.Println(name, lastName, age)

	name, lastName = "Victor", "Barbosa"
	fmt.Println(name, lastName, age)

	var height = 1.80
	fmt.Println(height)

	// Short variable declaration
	// only inside functions
	weight := 65.5
	fmt.Println(weight)
}
