package main

import "fmt"

func main() {
	var bilangan int64
	var angka1, angka2, angka3, angka4 int64

	fmt.Print("Input bilangan 4 digit positif: ")
	fmt.Scan(&bilangan)

	if bilangan >= 1000 && bilangan <= 9999 {

		angka1 = bilangan / 1000
		angka2 = (bilangan / 100) % 10
		angka3 = (bilangan / 10) % 10
		angka4 = bilangan % 10

		if angka1 < angka2 && angka2 < angka3 && angka3 < angka4 {
			fmt.Println("Terurut membesar")
		} else if angka1 >= angka2 && angka2 >= angka3 && angka3 >= angka4 {
			fmt.Println("Terurut mengecil")
		} else {
			fmt.Println("Tidak terurut")
		}
	} else {
		fmt.Println("inputkan 4 bilangan!")
	}
}
