package main

import (
	"fmt"
)

func main() {
	var umur int
	var kk bool

	fmt.Print("Masukkan usia: ")
	fmt.Scan(&umur)
	fmt.Print("Apakah memiliki kartu keluarga (true/false): ")
	fmt.Scan(&kk)

	if umur >= 17 && kk {
		fmt.Println("bisa membuat KTP")
	} else {
		fmt.Println("belum bisa membuat KTP")
	}
}
