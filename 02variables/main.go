package main

import "fmt"

// This link will give you brief overview about all the numeric types
// https://go.dev/ref/spec#Numeric_types

const LoginToken string = "dnjdfg"

// First capital letter of variable indicates its a public variables

func main() {
	var username string = "Divesh"
	fmt.Println(username)
	fmt.Printf("variables is of type: %T \n", username)

	var isValid bool = true
	fmt.Println(isValid)
	fmt.Printf("variables is of type: %T \n", isValid)

	var value int = 24
	fmt.Println(value)
	fmt.Printf("variables is of type: %T \n", value)

	// Default values and some aliases
	var anotherValue int
	fmt.Println(anotherValue)
	fmt.Printf("variables is of type: %T \n", anotherValue)

	//implicit type
	var website = "google.com"
	fmt.Println(website)

	//no var styles
	numberOfUsers := 30000
	// the above operator is called valurus operator it can only be use in a method only
	fmt.Println(numberOfUsers)
}
