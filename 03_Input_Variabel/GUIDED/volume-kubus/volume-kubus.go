package main

import (
	"fmt"
)

func main() {
	var sisi int

	fmt.Print("Masukkan panjang sisi kubus: ")
	fmt.Scan(&sisi)

	volume := sisi * sisi * sisi

	fmt.Printf("Volume kubus adalah: %0.d\n", volume)
}
