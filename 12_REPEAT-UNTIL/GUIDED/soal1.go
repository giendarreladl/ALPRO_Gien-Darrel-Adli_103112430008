package main

import "fmt"

func main() {
	var kata string
	var Banyaknya int

	fmt.Print("Masukkan kata: ")
	fmt.Scan(&kata)

	fmt.Print("Masukkan Banyaknya Perulangan: ")
	fmt.Scan(&Banyaknya)

	for i := 1; i <= Banyaknya; {
		fmt.Println(kata)
		i++
	}
}
