package helper

import "fmt"

var version = "1.0.0"
var Application = "Golang"

func sayGoodBye(name string) string {
	return "Good bye " + name
}

func SayHello(name string) string {
	//return "Hello " + name
	sayGoodBye("Yohan")
	fmt.Println(version)
}
