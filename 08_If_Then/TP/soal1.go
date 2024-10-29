package main

import "fmt"

func main() {
	var a int

	fmt.Print("Masukkan nilai anda: ")
	fmt.Scan(&a)

	if a >= 70 {
		fmt.Print("Anda Lulus")
	} else {
		fmt.Print("Anda Tidak lulus")
	}

}
