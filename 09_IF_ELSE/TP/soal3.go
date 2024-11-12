package main

import "fmt"

func main() {
	var umur int
	var kewarganegaraan string

	fmt.Print("Masukkan umur Anda: ")
	fmt.Scan(&umur)

	fmt.Print("Masukkan kewarganegaraan Anda: ")
	fmt.Scan(&kewarganegaraan)

	if umur >= 17 && (kewarganegaraan == "Indonesia" || kewarganegaraan == "indonesia") {
		fmt.Println("Anda bisa mengikuti pemilu")
	} else {
		fmt.Println("Anda tidak bisa mengikuti pemilu karena : ")
		if umur < 17 {
			fmt.Println("umur Anda belum mencapai 17 tahun")
		}
		if kewarganegaraan != "Indonesia" && kewarganegaraan != "indonesia" {
			fmt.Println("Anda bukan warga negara Indonesia")
		}
	}
}
