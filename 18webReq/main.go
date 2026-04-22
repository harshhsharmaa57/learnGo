package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://open.spotify.com/"

func main() {
	fmt.Println("Web request")

	res, err := http.Get(url)

	if err != nil {
		panic()
	}

	fmt.Printf("Response is of the type: %T\n", res)

	defer res.Body.Close()

	databytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)
}