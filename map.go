package main

import (
	"fmt"
)

func main() {
	person := map[string]string{
		"name":    "Alan",
		"address": "Surabaya",
	}

	//nambah key value

	person["status"] = "berhubungan gelap"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// jumlah isi dari map
	fmt.Println(len(person))

	// make(map[TypeKey]TypeValue) -> membuat map baru
	book := make(map[string]string)

	book["name"] = "Alan The Heart Compiler"
	book["author"] = "Muhammad Alan Nur"
	book["release date"] = "12-12-2021"

	fmt.Println(book)
	// delete(map, key) -> menghapus data di map dengan key
	delete(person, "name")
	fmt.Println(person)

}
