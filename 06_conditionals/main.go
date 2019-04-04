package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello World!")
	x := 2
	y := 5

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Println("woooow")
	}

	if 1 > 0 {
		//blah
	} else if 1 == 0 {
		//something else
	}

	// init a variable in the if statement
	if v := math.Pow(2, 3); v < 10 {
		fmt.Println(v)
	}

	//switch
	var test = "blue"
	switch test {
	case "red":
		fmt.Println("red!")
	case "blue":
		fmt.Println("Blue!")
	default:
		fmt.Println("default")
	}

}
