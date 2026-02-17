package main

import (
	"fmt"
)

var movies = []string{
	"O Poderoso Chefão",
	"Titanic",
	"Matrix",
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	fmt.Println(arr)
	fmt.Println(slice)

	newSlice := []int{1, 2, 3}
	fmt.Println(newSlice)
	fmt.Println(newSlice, len(newSlice), cap(newSlice))

	otherSlice := []int{}
	fmt.Println(otherSlice)

	resultsFromApi := []int{0, 1, 2}
	//var filmes []string

	matrix := [][]int{}
	fmt.Println(matrix)

	filmes := make([]string, 0, 3)
	for _, id := range resultsFromApi {
		filme := movies[id]
		fmt.Println(len(filmes))
		filmes = append(filmes, filme)
		fmt.Println(len(filmes))
	}

	fmt.Println(filmes)

	array := [5]int{1, 2, 3, 4, 5}
	mySlice := array[:2:2]
	fmt.Println(mySlice, len(mySlice), cap(mySlice))
}
