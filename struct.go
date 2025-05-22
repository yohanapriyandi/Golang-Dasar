package main

import "fmt"

type Customer struct {
	Name, Email, Address string
	Age                  int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello ", name, "My name is ", customer.Name)
}

func main() {
	var yohan Customer
	fmt.Println(yohan)

	yohan.Name = "Yohan Apriyandi"
	yohan.Email = "yohanapriyandi89@gmail.com"
	yohan.Address = "Gibug Cigadung"
	yohan.Age = 36

	fmt.Println(yohan)

	Fathar := Customer{
		Name:    "Fathar Dhabit Adz-Dzaky",
		Email:   "fda@gmail.com",
		Address: "Lingk. Gibug Cigadung",
		Age:     4,
	}

	fmt.Println(Fathar)

	Fariq := Customer{
		"Fariq Taqy Abqary",
		"fta@gmail.com",
		"Lingk. Gibug Cigadung",
		2,
	}

	fmt.Println(Fariq)
	
	Fathar.sayHello("Taqi")

}
