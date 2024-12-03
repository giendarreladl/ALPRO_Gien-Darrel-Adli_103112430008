package main

import "fmt"

func main() {
	var token string
	attempts := 0

	for {
		fmt.Print("Masukkan password: ")
		fmt.Scan(&token)

		if token == "12345abcde" {
			fmt.Println("Selamat, Anda berhasil login!")
			break
		} else {
			attempts++
			if attempts >= 3 {
				fmt.Println("Maaf, Anda tidak bisa mencoba lagi.")
				break
			}
			fmt.Println("Password salah, coba lagi.")
		}
	}
}
