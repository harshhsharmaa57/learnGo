package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "this needs to go in a file "

	file, err := os.Create("./file.txt")

	if err != nil {
		panic(err)
	} 

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("length is: ", length)
	defer file.Close()

	readFile("./file.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("text data inside file is in \n", string(databyte))
}

func checkNilErr (err error) {
	if err != nil {
		panic(err)
	} 
} 