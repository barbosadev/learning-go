package main

import "fmt"

func main() {
	m := make(map[string]string)

	otherMap := map[string]string{
		"Joaquim": "Pedro",
		"Pedro":   "Pessoa",
	}

	fmt.Println(m)
	fmt.Println(otherMap)

	sliceMap := map[string][]int{
		"Pedro":   {1, 2, 3},
		"Joaquim": {4, 5, 6},
	}
	delete(sliceMap, "Pedro")
	fmt.Println(sliceMap)

	clear(sliceMap)
	fmt.Println(sliceMap)

	oneMoreMap := make(map[string]string)
	oneMoreMap["Joaquim"] = "Pedro"

	value, ok := oneMoreMap["Joaquim"]

	fmt.Println(value, ok)
	fmt.Println(oneMoreMap)

	for k, v := range oneMoreMap {
		fmt.Println(k, v)
	}
}
