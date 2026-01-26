package main

import "fmt"

func main() {
	x := 10
	y := &x // y is a pointer to x
	fmt.Println(y, *y)
	modifyPointer(y)
	fmt.Println(y, *y)
}

func modifyPointer(p *int) {
	*p = 20
}
