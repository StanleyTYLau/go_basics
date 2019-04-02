package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	//define map
	emails := make(map[string]string)

	//assign key values
	emails["Bob"] = "bob@email.com"
	emails["Pog"] = "pb@email.com"
	emails["Foop"] = "Fooop@email.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Foop"])

	//delete
	delete(emails, "Pog")
	fmt.Println(emails)

	//declare and assign
	moreEmails := map[string]string{"bob": "bobby@email.com", "Pauuu": "PP@email.com"}
	fmt.Println("more emails: ", moreEmails)
}
