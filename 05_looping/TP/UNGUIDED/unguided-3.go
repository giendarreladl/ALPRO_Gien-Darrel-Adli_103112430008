package main

import "fmt"

func main() {
	var base, pangkat, hasil int

	fmt.Print("Masukkan bilangan pertama (basis): ")
	fmt.Scan(&base)

	fmt.Print("Masukkan bilangan kedua (pangkat): ")
	fmt.Scan(&pangkat)

	hasil = 1

	for i := 0; i < pangkat; i++ {
		hasil *= base
	}

	fmt.Println("Hasil dari", base, "pangkat", pangkat, "adalah: ", hasil)
}
