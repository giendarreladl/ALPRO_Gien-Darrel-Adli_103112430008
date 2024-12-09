package main

import "fmt"

func main() {
	var angka int

	for selesai := false; !selesai; {
		fmt.Print("masukkan angka :")
		fmt.Scan(&angka)

		selesai = (angka > 0)
	}

	fmt.Println(angka, "adalah bilangan positif")
}
