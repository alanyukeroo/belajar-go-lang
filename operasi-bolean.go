package main

import "fmt"

// !, &&, ||

func main() {

	var minAbsen = 10
	var minNilai = 75

	var jumlahAbsenku = 9
	var NilaiKu = 80

	var lulus bool = jumlahAbsenku >= minAbsen && NilaiKu >= minNilai

	fmt.Println(lulus)
}
