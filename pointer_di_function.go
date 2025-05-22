package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
	address.City = "Kuningan"
}

func main() {
	address := Address{}
	ChangeAddressToIndonesia(&address)

	fmt.Println(address)
}
