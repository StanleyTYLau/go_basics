package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	a := 10
	b := &a

	fmt.Println(a, b) //b prints out a memeory address - b is pointing at where a is stored
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b) //*int which is a pointer

	//reading a value from a pointer
	fmt.Println("value of b:", *b)

	//change value of pointer
	*b = 99
	fmt.Println("value of a:", a)

	//simply, pointer can be used to improve efficiency by passing memory location rather than a copy.

}
