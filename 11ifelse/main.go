package main

import "fmt"

func main() {
	fmt.Println("If else in go lang")
	loginCount := 23
	var result string

	if loginCount<13 {
		result = "regular access"
	} else if loginCount >13 {
		result = "watchout"
	} else {
		result = "13 login count"
	}
	fmt.Println(result)

	if 9%2==0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")

	}

	if num:=3; num<10 {
		fmt.Println("Num less than 10")
	} else {
		fmt.Println("Number not less than 10")
	}

	
}