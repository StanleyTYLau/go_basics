package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	ids := []int{12, 33, 44, 66, 2, 45, 99}

	//loop thru ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//loop not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//Range with map
	moreEmails := map[string]string{"bob": "bobby@email.com", "Pauuu": "PP@email.com"}

	for key, value := range moreEmails {
		//fmt.Println(key, value)
		fmt.Printf("%s: %s\n", key, value)

	}
}
