package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Muhammad"
	names[1] = "Alan"
	names[2] = "Nur"

	fmt.Println(names)

	var nomor [len(names)]int

	nomor[0] = 1
	nomor[1] = 2
	nomor[2] = 3

	fmt.Println(nomor)

	//multple declaration array

	var values = [3]int{
		90,
		80,
		95,
	}

	fmt.Println(values)

}
