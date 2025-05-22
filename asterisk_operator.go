package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// Asterisk Operator
	address1 := Address{"Kuningan", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address2.City = "Jakarta"
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println("==============================================================")
	//address2 = &Address{"Jakarta", "DKI jakarta", "Indonesia"}
	*address2 = Address{"Jakarta", "DKI jakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

}
