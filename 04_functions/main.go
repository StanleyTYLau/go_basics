package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func sum(a int, b int) int {
	//func sum(a , b int) int { >>> also works
	var answer int = a + b
	return answer
}

func main() {
	fmt.Println(greeting("Bob"))
	fmt.Println(sum(1, 2))

}
