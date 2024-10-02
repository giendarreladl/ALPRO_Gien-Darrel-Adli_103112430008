package main

import (
	"fmt"
)

func main() {
	var exchange float64 = 15000

	var rupiah float64

	fmt.Print("Masukkan jumlah rupiah: ")
	fmt.Scan(&rupiah)

	dolar := rupiah / exchange

	fmt.Printf("%.2f rupiah adalah %.2f dolar\n", rupiah, dolar)
}
