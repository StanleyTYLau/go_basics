package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	fmt.Println("Hello World!")

	sum := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}

	// Go functions may be closures. A closure is a function
	// value that references variables from outside its body.
	// The function may access and assign to the referenced variables;
	// in this sense the function is "bound" to the variables.
}
