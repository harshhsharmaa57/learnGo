package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://github.com/harshhsharmaa57?tab=repositories"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	
	qparams := result.Query()
	fmt.Printf("The Type of query params are: %T\n", qparams)
	
	fmt.Println(qparams["tab"])

	for _, val := range qparams {
		fmt.Println("Params is: ", val)
	}


	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "github.com",
		Path: "/harshhsharmaa57",
		RawQuery: "tab=repositories",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}