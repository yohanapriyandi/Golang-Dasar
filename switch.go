package main

import "fmt"

func main() {
	name := "Fathar"

	switch name {
	case "Yohan":
		fmt.Println("Hello Yohan")
	case "Fathar":
		fmt.Println("Hello Fathar")
	case "Taqy":
		fmt.Println("Hello Taqy")
	default:
		fmt.Println("Hello kamu siapa yah??!")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama sudah sesuai")
	case false:
		fmt.Println("nama terlalu panjang")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")

	}
}
