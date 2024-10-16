package main

import (
	"fmt"
)

func main() {
	var n int
	var alas, tinggi float32

	fmt.Print("Masukan Jumlah Segitiga Yang Ingin Dihitung : ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Println("segitiga", i)
		fmt.Print("Masukkan alas: ")
		fmt.Scan(&alas)
		fmt.Print("Masukkan tinggi: ")
		fmt.Scan(&tinggi)

		luas := 0.5 * alas * tinggi
		fmt.Println("Luas Segitiga", i, "adalah :", luas)
	}
}
