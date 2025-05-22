package main

import "fmt"

func main() {
	var names [5]string

	names[0] = "Fathar"
	names[1] = "Dhabit"
	names[2] = "Adz-dzaky"
	names[3] = "Yohan"
	names[4] = "Apriyandi"

	fmt.Println(names[2])
	fmt.Println(len(names))
	fmt.Println("========================================")

	var buah = [...]int{
		10, 20,
	}

	fmt.Println(len(buah))
	fmt.Println(buah)

}
