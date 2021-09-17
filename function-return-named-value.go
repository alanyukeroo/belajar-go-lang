package main

import "fmt"

func getFullName() (firstName, midName, lastName string) {
	firstName = "Muhammad"
	midName = "Alan"
	lastName = "Nur"
	return
}

func main() {
	a, b, c := getFullName()
	fmt.Println(a, b, c)
}
