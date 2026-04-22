package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")


	var fruitList = []string{"Apple","Tomato","Peach"}

	fmt.Printf("Type of fruitList is: %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	hightScore := make([]int, 4)
	hightScore[0] = 1
	hightScore[1] = 12
	hightScore[2] = 13
	hightScore[3] = 14

	
	fmt.Println(hightScore)

	hightScore = append(hightScore, 16,15,10)

	fmt.Println(hightScore)


	fmt.Println(sort.IntsAreSorted(hightScore))

	sort.Ints(hightScore)

	fmt.Println(hightScore)

	// how to remove a value from slices based on index

	var courses = []string{"javascript","java","ruby","swift","python","c"}
	fmt.Println(courses)

	var index int =2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}