package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	//long FOR loops
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
		//i += 1
	}

	//short
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//fizzbuzz
	//count 1 - 100, multiples of 3 = fizz, multiples of 5 = buzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("Fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

	}
}
