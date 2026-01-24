package main

import (
	"fmt"
	"names/pacote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(pacote.PublicName)
	fmt.Println(pacote.GetName())
}
