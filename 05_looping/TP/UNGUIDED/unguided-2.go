package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var r, t, volume float64

	// Meminta jumlah kerucut
	fmt.Print("Masukkan jumlah kerucut: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {

		fmt.Print("Masukkan jari-jari kerucut ke-", i, ": ")
		fmt.Scan(&r)

		fmt.Print("Masukkan tinggi kerucut ke-", i, ": ")
		fmt.Scan(&t)

		volume = (1.0 / 3.0) * math.Pi * r * r * t

		fmt.Printf("Volume kerucut ke-%d: %.5f\n", i, volume)
	}
}
