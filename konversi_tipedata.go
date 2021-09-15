package main

import "fmt"

func main() {

	var nilai32 int32 = 128
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	//overflow karena min - max int8 adalah -128 - 127. sehingga jika nilai melebihi batas, dia akan balik ke awal (overflow int)
	fmt.Println(nilai8)

	//konversi byte ke string
	var name = "Muhammad Alan Nur"
	var firstName = name[:8]
	var tipe_byte = name[0]
	fmt.Println(firstName)
	fmt.Println(string(tipe_byte))

}
