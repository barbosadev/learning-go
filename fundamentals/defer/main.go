package main

import "fmt"

func main() {
	doDefer()

	defer fmt.Println("Wake up!")
	defer september()
	fmt.Println("sleeping until september...")
}

func september() {
	fmt.Println("September...")
}

// lifo (last in, first out) order stack
func doDefer() {
	fmt.Println("before")
	defer fmt.Println("Print 1")
	defer fmt.Println("Print 2")
	defer fmt.Println("Print 3")
	fmt.Println("before too")
}
