package main

import "fmt"

func main() {
	type Married bool
	type iniString string

	var MarriedStatus Married = true
	var firstName iniString = "Alan Tampan"
	fmt.Println(MarriedStatus)
	fmt.Println(firstName)
}
