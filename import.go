package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Yohan")
	fmt.Println(result)

	fmt.Println(helper.Application)
	fmt.Println(helper.sayGoodBye) //tidak bisa dikases dari luar package
	fmt.Println(helper.version)    //tidak bisa dikases dari luar package
}
