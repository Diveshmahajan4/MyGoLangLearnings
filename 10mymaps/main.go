package main

import "fmt"

func main() {
	lang := make(map[string]string)

	lang["JS"] = "javascript"
	lang["TS"] = "typescript"
	lang["py"] = "python"
	lang["rb"] = "ruby"

	fmt.Println(lang)
	fmt.Println(lang["JS"])

	//loops

	for key, value := range lang {
		fmt.Println(key, " : ", value)
	}

	//for only values // comma ok syntax
	for _, value := range lang {
		fmt.Println(value)
	}

}
