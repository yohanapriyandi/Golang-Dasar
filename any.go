package main

import "fmt"

func Ups() any {
	return string("interface kosong")
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}
