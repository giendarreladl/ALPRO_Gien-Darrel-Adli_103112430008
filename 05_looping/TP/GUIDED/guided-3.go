package main

import (
	"fmt"
)

func main() {
	var inp1, inp2 int
	var hasil int = 0

	fmt.Print("Masukkan bilangan bulat pertama: ")
	fmt.Scan(&inp1)
	fmt.Print("Masukkan bilangan bulat kedua: ")
	fmt.Scan(&inp2)

	for i := 0; i < inp2; i++ {
		hasil += inp1
	}

	fmt.Println("Hasil perkalian", inp1, "x", inp2, "=", hasil)
}
