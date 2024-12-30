// Soal Nomor 7 (Sport Club)

//Pseudo Soal 8:

// =======program SportClub======

/*kamus
    s, minggu : integer
algoritma
    input(s)
    minggu = s div 7
    if s mod 7 == 0 then
        output("Minggu ke-", minggu)
    else
        output("Minggu ke-", minggu + 1)
    endif
endprogram*/

// Code Soal No 8:

package main

import (
	"fmt"
)

func main() {
	var s, minggu int

	fmt.Print("Masukkan jumlah hari: ")
	fmt.Scan(&s)

	minggu = s / 7
	if s%7 == 0 {
		fmt.Print("Minggu ke- ", minggu)
	} else {
		fmt.Print("Minggu ke- ", minggu+1)
	}
}
