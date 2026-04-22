package main

import "fmt"

func main() {
	fmt.Println("Structs in go lang")

	//no inheritance in go lang: no super or parent
	harsh := User{"Harsh", "harsh@gmail.com", true, 20}

	fmt.Println(harsh)

	fmt.Printf("Harsh details are : %+v\n", harsh)

	fmt.Printf("Harsh email: %v and Age is %v", harsh.Email, harsh.Age)

	harsh.GetStatus()
	harsh.NewMail()
	fmt.Printf("Harsh email: %v and Age is %v", harsh.Email, harsh.Age)


}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email) 
}