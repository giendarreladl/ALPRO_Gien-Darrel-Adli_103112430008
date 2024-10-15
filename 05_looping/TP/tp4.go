package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Mengatur seed agar angka acak berubah setiap kali program dijalankan

	// Program memilih angka acak antara 1 hingga 100
	target := rand.Intn(100) + 1
	var guess int

	// Memberikan 5 kesempatan untuk menebak
	for attempts := 1; attempts <= 5; attempts++ {
		fmt.Printf("Tebakan %d: Masukkan angka antara 1 hingga 100: ", attempts)
		fmt.Scan(&guess)

		if guess < target {
			fmt.Println("Terlalu kecil!")
		} else if guess > target {
			fmt.Println("Terlalu besar!")
		} else {
			fmt.Println("Selamat! Tebakanmu benar!")
			break
		}

		// Jika percobaan terakhir dan tebakan masih salah
		if attempts == 5 {
			fmt.Println("Kamu telah mencoba 5 kali. Angka yang benar adalah :", target)
		}
	}
}
