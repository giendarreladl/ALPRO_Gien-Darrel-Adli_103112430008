package main

import "fmt"

func main() {
	const correctPassword = "12345"
	const maxAttempts = 3

	var input string

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		fmt.Printf("Masukkan password (Percobaan %d dari %d): ", attempts, maxAttempts)
		fmt.Scanln(&input)

		if input == correctPassword {
			fmt.Println("Login berhasil!")
			return
		} else {
			fmt.Println("Password salah.")
		}
	}

	fmt.Println("Login ditolak. Anda telah mencoba 3 kali.")
}
