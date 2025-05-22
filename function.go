package main

import "fmt"

func sayHello1(fisrtName string, lastName string, age int64) {
	fmt.Println("Hello", fisrtName, lastName, age)
}

func main() {
	sayHello1("Yohan", "Apriyandi", 36)
	sayHello1("Fithriya", "Nabilah", 28)
	sayHello1("Fathar Dhabit", "Adz-dzaky", 4)
}
