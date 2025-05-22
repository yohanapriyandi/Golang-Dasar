package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// Pass by value
	//address1 := Address{"Kuningan", "Jawa Barat", "Indonesia"}
	//address2 := address1

	// Pass by reference
	address1 := Address{"Kuningan", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address2.City = "Jakarta"
	address2.Province = "DKI Jakarta"
	fmt.Println(address1)
	fmt.Println(address2)
}
