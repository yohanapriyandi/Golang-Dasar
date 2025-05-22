package main

import "fmt"

func main() {
	name := "kim yeo Hoon"

	if name == "Yohan" {
		fmt.Println("Helo Yohan")
	} else if name == "Fathar" {
		fmt.Println("Hai Fathar")
	} else if name == "Taqy" {
		fmt.Println("Hai aku Fariq Taqy Abqary ")
	} else {
		fmt.Println("Hai kamu ini siapa!!!")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
