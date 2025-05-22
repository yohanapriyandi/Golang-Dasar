package main

import "fmt"

func random() any {
	return false
}

func main() {
	var result = random()
	//var resultString = result.(string)
	//fmt.Println(resultString)
	//
	//var resultInt = result.(int)
	//fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	default:
		fmt.Println("Unknown", value)

	}
}
