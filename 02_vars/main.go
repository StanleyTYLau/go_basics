package main

import "fmt"

func main() {
	//Main Types
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr - unsigned integer
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	//using var
	var name string = "wow a string!"
	var name2 = "string type is inferred."
	var number int32 = 123
	const pi = 3.14159

	//shorthand method can only be used inside a function
	happy := "cool!"
	guy, email := "Bob", "bob@email.com"

	fmt.Println(name, name2, number, pi, happy, guy, email)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", number)
	fmt.Printf("%T\n", pi)
}
