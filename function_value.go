package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	contoh1 := getGoodBye
	contoh2 := getGoodBye

	fmt.Println(contoh1("Fathar"))
	fmt.Println(contoh2("Taqy"))

}
