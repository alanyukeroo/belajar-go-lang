package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World!")
}

func nama_lengkap(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func luas_persegi(p int, l int) int {
	var hasil = p * l
	return hasil
}

func hasil_list(x int, y int) []int {
	a := make([]int, x)
	for i := 0; i < x; i++ {
		a[i] = y
	}
	return a

}

func main() {
	sayHello()
	nama_lengkap("alan", "tampan")
	fmt.Println(luas_persegi(3, 3))
	fmt.Println(hasil_list(4, 5))
}
