package main

import "fmt"

func main() {
	//counter := 5
	//
	//for counter <= 10 {
	//	fmt.Println("Perulangan ke ", counter)
	//	counter++
	//}

	for counter := 5; counter <= 10; counter++ {
		fmt.Println("Perualangan ke ", counter)
	}

	fmt.Println("Selesai")

	names := []string{"Fathar", "Dhabit", "Adz-dzaky"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("Index ", index, "=", name)
	}
	fmt.Println("================================================")
	for _, name := range names {
		fmt.Println(name)
	}
}
