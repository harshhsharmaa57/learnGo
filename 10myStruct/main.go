package main

import "fmt"

func main() {
	fmt.Println("Structs in go lang")

	//no inheritance in go lang: no super or parent
	harsh := User{"Harsh", "harsh@gmail.com", true, 20}

	fmt.Println(harsh)

	fmt.Printf("Harsh details are : %+v\n", harsh)

	fmt.Printf("Harsh email: %v and Age is %v", harsh.Email, harsh.Age)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}