package main

import "fmt"

func main() {
	counter := 0

	//fitur closure di go
	increment := func() {
		fmt.Println("Increment ")
		counter++
	}

	increment()
	increment()
	increment()
	increment()
	increment()
	fmt.Println(counter)
}
