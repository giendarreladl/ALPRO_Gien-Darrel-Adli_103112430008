package main

import (
	"fmt"
)

func main() {
	const usernameBenar = "Admin"
	const passwordBenar = "Admin"

	var jumlahGagal int

	for {
		var username, password string
		fmt.Println("Masukkan nama pengguna dan kata sandi (pisahkan dengan spasi):")
		_, err := fmt.Scanln(&username, &password)

		if err != nil {
			fmt.Println("Input tidak valid. Harap masukkan nama pengguna dan kata sandi dengan format 'nama_pengguna kata_sandi'")
			continue
		}

		if username == usernameBenar && password == passwordBenar {
			fmt.Printf("%d percobaan gagal login\n", jumlahGagal)
			break
		} else {
			jumlahGagal++
			fmt.Println("Nama pengguna atau kata sandi salah, coba lagi.")
		}
	}
}
