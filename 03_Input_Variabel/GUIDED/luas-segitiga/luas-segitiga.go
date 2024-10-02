package main

import (
	"fmt"
)

func main() {
	var alas float64
	var tinggi float64

	fmt.Print("Masukkan panjang alas segitiga: ")
	fmt.Scan(&alas)
	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scan(&tinggi)

	luas := 0.5 * alas * tinggi

	fmt.Printf("Volume luas segitiga adalah: %0.f\n", luas)
}
