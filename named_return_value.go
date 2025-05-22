package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Fathar"
	middleName = "Dhabit"
	lastName = "Adz-Dzaky"

	return firstName, middleName, lastName

}

func main() {
	f, m, l := getCompleteName()
	fmt.Println(f, m, l)
}
