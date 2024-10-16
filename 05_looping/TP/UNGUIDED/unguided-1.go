package main

import "fmt"

func main() {
	var input, total int

	// Meminta masukan dari pengguna
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&input)

	// Menghitung jumlah bilangan dari 1 sampai input
	for i := 1; i <= input; i++ {
		total += i
	}

	// Menampilkan hasil keluaran
	fmt.Println("Keluaran: ", total)
}
