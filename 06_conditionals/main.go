package main

import "fmt"

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
