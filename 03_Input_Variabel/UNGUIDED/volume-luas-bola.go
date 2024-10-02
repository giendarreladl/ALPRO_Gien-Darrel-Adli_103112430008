package main

import (
	"fmt"
)

func main() {

	var jari float32
	var volume float32
	var luas float32

	fmt.Print("Masukkan nilai jari-jari bola: ")
	fmt.Scan(&jari)

	volume = 3.14 * jari * jari * jari / 3 * 4
	luas = 4 * 3.14 * jari * jari

	fmt.Println("Volume bola dengan jari-jari", jari, "adalah:", volume, "dan luas kulit adalah:", luas)
}
