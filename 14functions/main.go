package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	
	greeterTwo()


	result := adder(3,5)

	fmt.Println("result is: ", result)

	proRes, myMessage = proAdder(1,3,4,5,6,67,78,8)
	fmt.Println("pro result is: ", proRes)
	fmt.Println("Pro Message is: ", myMessage)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int, string{
	total := 0
	for _,val := range values {
		total += val
	}
	return total, "HI"
}

func greeter() {
	fmt.Println("namastey from goland")
}

func greeterTwo() {
		fmt.Println("Another method")
} 