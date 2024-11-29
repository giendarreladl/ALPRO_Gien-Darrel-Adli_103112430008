package main

import "fmt"

func main() {
	var (
		totalBelanja, harga float64
		namaBarang, pilihan string
	)

	fmt.Println("=== Aplikasi Kasir Sederhana ===")

	for {
		fmt.Print("Masukkan nama barang: ")
		fmt.Scanln(&namaBarang)

		fmt.Print("Masukkan harga barang: ")
		fmt.Scanln(&harga)

		totalBelanja += harga
		fmt.Printf("Barang '%s' dengan harga %.2f berhasil ditambahkan.\n", namaBarang, harga)

		fmt.Print("Tambah barang lagi? (y/n): ")
		fmt.Scanln(&pilihan)

		if pilihan == "n" || pilihan == "N" {
			fmt.Println("=== Transaksi Selesai ===")
			fmt.Printf("Total belanja: Rp%.2f\n", totalBelanja)
			break
		}
	}
}
