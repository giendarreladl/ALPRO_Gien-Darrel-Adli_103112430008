package main

import "fmt"

func main() {
	var (
		BB, TB, KG, Hasil float64
	)

	fmt.Print("Masukan BB :")
	fmt.Scan(&BB)
	fmt.Print("Masukan TB :")
	fmt.Scan(&KG)

	TB = KG * KG
	Hasil = BB / TB

	fmt.Printf("Hasil Nilai BMI %.2f", Hasil)

}
