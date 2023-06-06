package main

import (
	"fmt"
	"sort"
)

func main() {
	// var fruits = []string{"apple", "mango", "banana"}
	// fmt.Printf("type of fruit list is %T \n", fruits)

	// fruits = append(fruits, "guava", "strawberry")
	// fmt.Println(fruits)

	// fruits = append(fruits[1:3])
	// fmt.Println(fruits)

	scores := make([]int, 4)

	scores[0] = 12
	scores[1] = 56
	scores[2] = 34
	scores[3] = 16

	scores = append(scores, 25, 78)
	fmt.Println(scores)

	fmt.Println(sort.IntsAreSorted(scores))
	sort.Ints(scores)
	fmt.Println(scores)

	//how to remove value form slices
	var langs = []string{"java", "python", "cpp", "c", "go", "ruby"}
	fmt.Println(langs)
	// we have to remove index 2
	var num int = 2
	langs = append(langs[:num], langs[num+1:]...)
	fmt.Println(langs)
}
