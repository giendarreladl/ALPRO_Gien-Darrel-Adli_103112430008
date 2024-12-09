package main

import "fmt"

func main() {
	var Bilangan1 int
	var Bilangan2 int

	fmt.Print("Masukkan x: ")
	fmt.Scan(&Bilangan1)

	fmt.Print("Masukkan y: ")
	fmt.Scan(&Bilangan2)

	for selesai := false; !selesai; {
		Bilangan1 = Bilangan1 - Bilangan2
		fmt.Println(Bilangan1)
		selesai = Bilangan1 <= 0
	}

	fmt.Println(Bilangan1 == 0)
}
