package delivery

import (
	"fmt"
	"strings"
)

func MainMenu() {
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Warung Makan Bahari")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("1. Tambahkan Menu Baru")
	fmt.Println("2. Daftar Menu")
	fmt.Println("3. Cari Menu")
	fmt.Println("4. Keluar")
	fmt.Println("Pilih Menu (1-4)")
}
