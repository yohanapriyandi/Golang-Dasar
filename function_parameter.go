package main

import "fmt"

func sayHello(firstName string, lastName string, age int64) {
	fmt.Println("Hello", firstName, lastName, "Umur:", age)
}

func main() {
	sayHello("Yohan", "Apriyandi", 36)
}
