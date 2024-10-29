package main

import "fmt"

func main() {
	var a int

	fmt.Print("Masukkan angka anda: ")
	fmt.Scan(&a)

	b := a % 2

	if b == 0 {
		fmt.Print("Angka adalah genap")
	} else {
		fmt.Print("Angka adalah ganjil")
	}

}
