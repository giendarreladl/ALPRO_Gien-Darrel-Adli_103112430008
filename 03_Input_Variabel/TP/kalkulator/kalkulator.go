package main

import (
	"fmt"
)

func main() {

	var inp float32
	var angka1 float32
	var angka2 float32

	fmt.Println("Pilihlah Aritmatika berikut!:")
	fmt.Println("1. Penjumlahan")
	fmt.Println("2. Pengurangan")
	fmt.Println("3. Pembagian")
	fmt.Println("4. Perkalian")
	fmt.Scanln(&inp)

	if inp == 1 {
		fmt.Println("Masukan angkaka pertama :")
		fmt.Scanln(&angka1)
		fmt.Println("Masukan angkaka kedua :")
		fmt.Scanln(&angka2)
		fmt.Println("Hasil Penjumlahan adalah :", angka1+angka2)
	} else if inp == 2 {
		fmt.Println("Masukan angkaka pertama")
		fmt.Scanln(&angka1)
		fmt.Println("Masukan angkaka kedua")
		fmt.Scanln(&angka2)
		fmt.Println("Hasil Pengurangkaan adalah", angka1-angka2)
	} else if inp == 3 {
		fmt.Println("Masukan angkaka pertama")
		fmt.Scanln(&angka1)
		fmt.Println("Masukan angkaka kedua")
		fmt.Scanln(&angka2)
		fmt.Println("Hasil pembagian adalah", angka1/angka2)
	} else if inp == 4 {
		fmt.Println("Masukan angkaka pertama")
		fmt.Scanln(&angka1)
		fmt.Println("Masukan angkaka kedua")
		fmt.Scanln(&angka2)
		fmt.Println("Hasil Perkalian adalah", angka1*angka2)
	} else {
		fmt.Print("Tidak ada pilihan!")
	}
}
