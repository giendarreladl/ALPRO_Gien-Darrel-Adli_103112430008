package main

import (
	"fmt"
)

func main() {
	var kantong1, kantong2 float64
	var totalBerat float64
	selesai := false

	for !selesai {
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kantong1, &kantong2)

		if kantong1 >= 9 || kantong2 >= 9 {
			fmt.Println("Proses selesai.")
			break
		}

		selisih := kantong1 - kantong2
		if selisih < 0 {
			selisih = -selisih
		}
		fmt.Printf("Sepeda motor Pak Andi akan oleng: %v\n", selisih > 9)

		totalBerat = kantong1 + kantong2
		fmt.Printf("Total berat kedua kantong: %.2f\n", totalBerat)

		if totalBerat > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		// Periksa jika salah satu kantong bernilai negatif
		if kantong1 < 0 || kantong2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}
	}
}
