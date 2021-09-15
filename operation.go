package main

import "fmt"

func main() {
	var a = 1
	var b = 2
	var c = 3

	fmt.Println(a + b + c)

	a++
	fmt.Println(a)

	a -= 3
	fmt.Println(a)

	a = -10
	fmt.Println(a)

}
