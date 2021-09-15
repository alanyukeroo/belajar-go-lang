package main

import "fmt"

func main() {
	var nama string

	nama = "Muhammad Alan Nur"
	fmt.Println(nama)

	nama = "Eh aku ganti nama jadi Alan tampan :p"
	fmt.Println(nama)

	// nama = true //error karena var tidak bisa diganti tipe datanya (str -> bool)
	// name = 123 //error karena var tidak bisa digantikan tipedatanya (str -> int)

	//variable dalam go juga dapat secara langsung mengidentifikasi tipe data tanpa melakukan deklarasi
	var namaLengkap = "Ini nama lengkapku haha"
	fmt.Println(namaLengkap)

	var age = 23
	fmt.Println(age)

	//alternatif dekralasi awal variable menggunakan simbol :=

	nama_ibu := "Momo"
	fmt.Println(nama_ibu)
	nama_ibu = "Mama"
	fmt.Println(nama_ibu)

	//deklarasi lebih dari satu variable
	var (
		firstName = "Alan"
		lastName  = "Nur"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}
