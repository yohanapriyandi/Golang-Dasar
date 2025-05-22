package main

import "fmt"

func getFullName() (string, string) {
	return "Yohan", "Apriyandi"
}

func main() {
	//fmt.Println("Memapilkan semua nilai")
	//firstName, lastName := getFullName()
	//fmt.Println(firstName, lastName)

	fmt.Println("Menampilkan sebagian nilai dari variabel saja")
	firstName, _ := getFullName()
	fmt.Println(firstName)
}
