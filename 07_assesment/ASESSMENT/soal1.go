package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan nilai x (lebih kecil atau sama dengan y): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	if x > y {
		fmt.Println("Nilai x harus lebih kecil atau sama dengan y!")
		return
	}

	total := 0
	for i := x; i <= y; i++ {
		total += i
	}

	fmt.Print("Hasil :", x, y, total)
}
