package main

import "fmt"

func main() {
	type NoKTP string

	var ktpYohan NoKTP = "11111111111"

	var contoh string = "2222222222"
	var contohktp NoKTP = NoKTP(contoh)

	fmt.Println(ktpYohan)
	fmt.Println(contohktp)
}
