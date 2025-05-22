package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	eko := Man{"Yohan"}
	eko.Married()

	fmt.Println(eko.Name)
}
