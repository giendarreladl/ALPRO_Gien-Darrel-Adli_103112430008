package main

import (
	"fmt"
)

func main() {
	var waktu int
	var harga int
	var kendaraan string
	fmt.Print("Masukkan Jenis kendaraan Motor/Mobil/Truk: ")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukkan lama durasi parkir (jam) : ")
	fmt.Scan(&waktu)

	switch {
	case kendaraan == "Motor" && waktu >= 1 && waktu <= 2:
		harga = 7000
	case kendaraan == "Motor" && waktu > 2:
		harga = 9000
	case kendaraan == "Mobil" && waktu >= 1 && waktu <= 2:
		harga = 15000
	case kendaraan == "Mobil" && waktu > 2:
		harga = 20000
	case kendaraan == "Truk" && waktu >= 1 && waktu <= 2:
		harga = 25000
	case kendaraan == "Truk" && waktu > 2:
		harga = 30000
	default:
		fmt.Print("Jenis kendaraan / durasi waktu parkir tidak valid")
	}
	fmt.Printf("Tarif Parkir : Rp. %d\n", harga)
}
