package main

import "fmt"

func main() {
	var (
		bilangan, angka1, angka2, angka3 int64
	)

	fmt.Print("input bilangan 3 digit positif :")
	fmt.Scan(&bilangan)

	angka1 = bilangan / 100
	angka2 = (bilangan % 100) / 10
	angka3 = bilangan % 10

	fmt.Println(angka1 <= angka2 && angka2 <= angka3)

}
