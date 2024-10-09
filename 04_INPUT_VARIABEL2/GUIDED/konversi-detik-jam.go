package main

import "fmt"

func main() {
	var (
		jam, menit, detik, masukan int64
	)

	fmt.Print("masukan detik :")
	fmt.Scan(&masukan)

	jam = masukan / 3600
	menit = masukan / 60
	detik = masukan

	fmt.Println("hasil jam: ", jam)
	fmt.Println("hasil menit: ", menit)
	fmt.Println("hasil detik: ", detik)

}
