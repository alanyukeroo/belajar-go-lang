package main

import "fmt"

func multipleAll(numbers ...int) int {
	var total int = 1

	for _, value := range numbers {
		total = total * value
	}
	return total
}

func main() {
	total := multipleAll(1, 2, 3, 4, 5, 6)
	fmt.Println(total)
}
