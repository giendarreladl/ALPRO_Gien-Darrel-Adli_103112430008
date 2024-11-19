package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Print("Masukkan usia Anda: ")
	fmt.Scanln(&age)

	switch {
	case age >= 0 && age <= 12:
		fmt.Println("Kategori usia: Anak-anak")
	case age >= 13 && age <= 17:
		fmt.Println("Kategori usia: Remaja")
	case age >= 18 && age <= 64:
		fmt.Println("Kategori usia: Dewasa")
	case age >= 65:
		fmt.Println("Kategori usia: Lansia")
	default:
		fmt.Println("Usia tidak valid")
	}
}
