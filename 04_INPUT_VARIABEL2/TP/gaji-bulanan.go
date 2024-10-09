package main

import "fmt"

func main() {
	var jamKerjaMingguan float64
	var GajiPerJam float64

	fmt.Print("Masukkan jumlah jam kerja per minggu: ")
	fmt.Scan(&jamKerjaMingguan)

	fmt.Print("Masukkan gaji per jam: ")
	fmt.Scan(&GajiPerJam)

	var jamNormal, jamLembur float64
	if jamKerjaMingguan > 40 {
		jamNormal = 40
		jamLembur = jamKerjaMingguan - 40
	} else {
		jamNormal = jamKerjaMingguan
		jamLembur = 0
	}

	gajiMingguan := (jamNormal * GajiPerJam) + (jamLembur * 1.5 * GajiPerJam)

	gajiBulanan := gajiMingguan * 4

	fmt.Println("Total gaji bulanan: ", int(gajiBulanan))
}
