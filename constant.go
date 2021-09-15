package main

import "fmt"

func main() {
	const firstName = "Muhammad"
	const lastName = "Alan Nur"

	//constant tidak masalah jika tidak dipanggil beda dengan variable yang harus digunakan / dipanggil
	//pada kasus ini saya hanya memanggil constant lastName
	fmt.Println(lastName)

	//bisa multiple declaration seperti variabel
	const (
		age  = 23
		city = "kudus"
	)
}
