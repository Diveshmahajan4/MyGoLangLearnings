package main

import "fmt"

func main() {
	// arrays

	var fruits [4]string

	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[3] = "mango"

	fmt.Println(fruits)
	fmt.Println(len(fruits))

	var veg = [5]string{"spinach", "tomato", "potato"}
	fmt.Println(veg)
}
