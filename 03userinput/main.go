package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to take input from user")
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating")
	input, _ := reader.ReadString('\n')

	fmt.Println("The rating is, ", input)
	fmt.Printf("the type of input is %T", input)
}