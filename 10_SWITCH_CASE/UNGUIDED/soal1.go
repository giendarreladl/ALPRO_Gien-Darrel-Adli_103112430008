package main

import (
	"fmt"
)

func main() {
	var ph float32

	fmt.Print("Masukan pH air dari 0-14: ")
	fmt.Scan(&ph)

	switch {
	case ph >= 6.5 && ph <= 8.6:
		fmt.Println("PH: ", ph, "Air layak minum ")
	case ph >= 0 && ph <= 14:
		fmt.Println("PH : ", ph, "Air tidak layak minum ")
	default:
		fmt.Print("Nilai ph tidak valid. Nilai ph harus antara 0 dan 14.")

	}
}
