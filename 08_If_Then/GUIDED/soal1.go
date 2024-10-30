package main

import "fmt"

func main() {
	var a, nilai int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&a)

	if a < 0 {
		nilai = -a
	} else {
		nilai = a
	}

	fmt.Print("nilai absolut dari ", a, " adalah : ", nilai)

}
