package main

import "fmt"

func getFullName() (string, string) {
	return "Muhammad", "Alan Nur"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}
