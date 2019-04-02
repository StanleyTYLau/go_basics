package main

import (
	"fmt"
	"strconv"
)

// define Person struct
type Person struct {
	firstName string
	lastName  string
	age       int

	// fisrtName, lastName string
	// age                 int
}

//Methods for struct goes outside
//Greeting method (value reciever - did not change values of the struct. Only did calculations or get some value)
func (p Person) greet() string {
	return "Hello, I am " + p.firstName + ". age: " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	fmt.Println("structs are similar to classes!")

	//initialize Person
	person1 := Person{firstName: "Bob", lastName: "Popper", age: 29}

	//alternative init
	person2 := Person{"Phoool", "Doop", 99}

	fmt.Println(person1, person2)

	person1.firstName = "Goober"
	fmt.Println(person1.firstName)
	fmt.Println(person1.greet())

	person1.hasBirthday()
	fmt.Println(person1.greet())

}
