// ===========PSEUDOCODE============

/* kamus
    siswa, mobil, remaining : integer
    kapasitas : constant integer
algoritma
    kapasitas = 7
    input(siswa)
    mobil = siswa div kapasitas
    remaining = siswa mod kapasitas
    if remaining > 0 then
        mobil = mobil + 1
        output(mobil - 1, "mobil penuh dan 1 mobil berisi", remaining, "orang")
    else
        output(mobil, "mobil penuh")
    endif
endprogram*/

// ===========Code============
package main

import (
	"fmt"
)

func main() {

	const kapasitas = 7

	var siswa int
	fmt.Println("Masukkan jumlah mahasiswa:")
	fmt.Scanf("%d", &siswa)

	mobil := siswa / kapasitas
	remaining := siswa % kapasitas

	if remaining > 0 {
		mobil++
		fmt.Printf("%d mobil penuh dan 1 mobil berisi %d orang\n", mobil-1, remaining)
	} else {
		fmt.Printf("%d mobil penuh\n", mobil)
	}
}
