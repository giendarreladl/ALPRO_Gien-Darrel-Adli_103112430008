package main

import (
	"fmt"
)

func main() {
	var harga, totalBelanja int

	for {
		fmt.Print("Masukkan harga barang (ketik 0 untuk selesai): ")
		fmt.Scan(&harga)

		if harga == 0 {
			break
		}
		totalBelanja += harga
	}

	fmt.Printf("Total belanja Anda: %d\n", totalBelanja)
}
