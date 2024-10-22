package main

import "fmt"

func main() {
	var jumlahBarang, totalPoin int

	fmt.Print("Masukkan jumlah barang yang dibeli: ")
	fmt.Scan(&jumlahBarang)

	// Menghitung poin menggunakan looping
	for i := 1; i <= jumlahBarang; i++ {
		if i <= 5 {
			totalPoin += 10 // Barang pertama sampai kelima memberikan 10 poin
		} else {
			totalPoin += 15 // Barang ke-6 dan seterusnya memberikan 15 poin
		}
	}

	fmt.Println("Total poin yang didapatkan:", totalPoin, "poin")
}
