package main

import "fmt"

func main() {
	var a = 20
	var b = 10
	var c = a*b - a + b

	fmt.Println(c)

	var i = 10

	i += 10 // explain -> i = i + 10, output 20
	fmt.Println(i)

	i += 5 // explain -> i = i + 5, output 25
	fmt.Println(i)

	i -= 3
	fmt.Println(i)

	i *= 5
	fmt.Println(i)

	// Urnary operation
	var j = 1

	j++
	fmt.Println(j)
	j++
	fmt.Println(j)

	j--
	fmt.Println(j)
	j--
	fmt.Println(j)

}
