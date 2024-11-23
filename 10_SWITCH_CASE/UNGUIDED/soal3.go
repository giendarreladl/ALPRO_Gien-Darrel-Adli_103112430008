package main

import (
	"fmt"
)

func main() {
	var bilangan int

	fmt.Println("Masukkan bilangan bulat:")
	fmt.Scanln(&bilangan)

	switch {
	case bilangan == 5:
		// Khusus bilangan 5
		hasil := bilangan + (bilangan + 1)
		fmt.Printf("Kategori: Bilangan Ganjil\n")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", bilangan, bilangan+1, hasil)
	case bilangan/10*10 == bilangan:
		// Kelipatan 10
		hasil := bilangan / 10
		fmt.Printf("Kategori: Bilangan Kelipatan 10\n")
		fmt.Printf("Hasil pembagian antara %d / 10 = %d\n", bilangan, hasil)

	case bilangan/5*5 == bilangan:
		// Kelipatan 5
		hasil := bilangan * bilangan
		fmt.Printf("Kategori: Bilangan Kelipatan 5\n")
		fmt.Printf("Hasil kuadrat dari %d ^ 2 = %d\n", bilangan, hasil)

	case bilangan/2*2 == bilangan:
		// Bilangan Genap
		hasil := bilangan * (bilangan + 1)
		fmt.Printf("Kategori: Bilangan Genap\n")
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n", bilangan, bilangan+1, hasil)

	default:
		// Bilangan Ganjil
		hasil := bilangan + (bilangan + 1)
		fmt.Printf("Kategori: Bilangan Ganjil\n")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", bilangan, bilangan+1, hasil)
	}
}
