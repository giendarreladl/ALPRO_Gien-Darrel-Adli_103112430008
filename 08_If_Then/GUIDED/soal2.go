package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)

	if bilangan < 1 {
		fmt.Print("Bukan Positif")
	} else {
		fmt.Print("Positif")
	}

}
