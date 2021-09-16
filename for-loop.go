package main

import "fmt"

func main() {

	//Kode program for
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	//For dengan statement
	var i int

	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for range -> digunakan untuk iterasi terhadap semua data collection
	// contoh data collection -> Array, Slice, dan Map

	var names = []string{"Alan", "Tampan", "Sekali"}

	for index, name := range names {
		fmt.Println(index, name)
	}

}
