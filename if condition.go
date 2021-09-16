package main

import "fmt"

func main() {
	var name string = "momomomomom"

	if name == "alan" {
		fmt.Println("Halo Alan!")
	} else if name == "momo" {
		fmt.Println("Halo Momo!")
	} else {
		fmt.Println("Eh kamu siapa?")
	}

	if lenght := len(name); lenght > 5 {
		fmt.Println("Nama Terlalu Dowo")
	} else {
		fmt.Println("Oke Nama Sudah Benar!")
	}

}
