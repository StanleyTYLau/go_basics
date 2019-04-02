package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	// Arrays
	var array1 [2]string
	array1[0] = "apple"
	array1[1] = "grape"
	fmt.Println(array1)

	//declare and assign
	var fruits = [2]string{"apples", "oranges"}
	//shorthand
	//fruits := [2]string{"apples", "oranges"}
	fmt.Println(fruits)

	//slices
	var fruits2 = []string{"a", "b", "c"}
	//fruits2 := []string{"a","b","c"}
	fmt.Println(fruits2)

	fmt.Println(len(fruits2))
	fmt.Println(fruits2[1:3])

	var newFruits = append(fruits2, "melon") //append only works on slices
	fmt.Println(newFruits)

}
