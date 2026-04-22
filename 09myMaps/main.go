package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["rb"] = "ruby"

	fmt.Println("List of all laguages: ", languages)
	fmt.Println("Js short for: ", languages["js"])

	delete(languages,"rb")

	fmt.Println("List of all languages : ", languages)


	//loops are interesting in go

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("For key v, value is v\n", value)
	}
}