package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println("The value of the pointer is", ptr)

	num := 10
	var ptr1 = &num
	fmt.Println("The value of the pointer is", ptr1)
	fmt.Println("The value of the pointer is", *ptr1)

	*ptr1 = *ptr1 + 2
	fmt.Println(ptr1)
}
