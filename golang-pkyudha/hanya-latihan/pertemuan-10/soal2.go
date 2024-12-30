package main

import (
	"fmt"
)

func main() {
	var nilai float64

	fmt.Print("Masukkan nilai TAK: ")
	fmt.Scan(&nilai)

	if nilai < 2.0 {
		fmt.Println("Poor")
	} else if nilai >= 2.0 && nilai <= 2.75 {
		fmt.Println("Fair")
	} else if nilai > 2.75 && nilai <= 3.0 {
		fmt.Println("Satisfactory")
	} else if nilai > 3.0 && nilai <= 3.5 {
		fmt.Println("Very Good")
	} else if nilai > 3.5 && nilai <= 4.0 {
		fmt.Println("Excellent")
	} else {
		fmt.Println("Nilai tidak valid")
	}
}
