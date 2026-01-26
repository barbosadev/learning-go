package main

import "fmt"

func main() {
	do(1)
	do(2)
	do(3)
	do(4)
	do(5)

	compare(1)
	compare(2)
	compare(3)
	compare(4)
	compare(5)
}

func compare(x int) {
	switch {
	case x == 1:
		fmt.Println("One")
	case x == 2:
		fmt.Println("Two")
	case x == 3 || x == 4:
		fmt.Println("Three or Four")
	default:
		fmt.Println("Other")
	}
}

func do(x int) {
	switch x {
	case 1:
		fmt.Println("One")
		break
		// break is optional here
	case 2:
		fmt.Println("Two")
		fallthrough
		// fallthrough forces the next case to run
	case 3, 4:
		fmt.Println("Three or Four")
	default:
		fmt.Println("Other")
	}
}
