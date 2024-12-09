package main

import "fmt"

func main() {
	var target, donasi, totalDonasi, jumlahDonatur int

	fmt.Print("Masukkan target donasi: ")
	fmt.Scanln(&target)

	for {
		fmt.Print("Masukkan jumlah donasi: ")
		fmt.Scanln(&donasi)

		totalDonasi += donasi
		jumlahDonatur++

		fmt.Println("Donatur", jumlahDonatur, "Menyumbang", donasi, ". Total terkumpul: ", totalDonasi)

		if totalDonasi >= target {
			break
		}
	}

	fmt.Println("Target tercapai! Total donasi:", totalDonasi, "dari", jumlahDonatur, "donatur.")
}
