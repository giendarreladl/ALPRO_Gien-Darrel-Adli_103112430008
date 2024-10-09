package main

import "fmt"

func main() {
	var (
		Harga, Diskon, Hasil float64
	)

	fmt.Print("Masukan total harga :")
	fmt.Scan(&Harga)
	fmt.Print("Masukan Diskon :")
	fmt.Scan(&Diskon)

	Hasil = Harga - (Harga * Diskon / 100)

	fmt.Println("total harga belanja adalah :", Hasil)

}
