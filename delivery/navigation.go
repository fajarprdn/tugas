package delivery

import (
	"fmt"
	"os"
	"warung_makan_bahari/delivery/util"
)

func ExitToMain() {
	var mainMenuConfirmation string
	fmt.Printf(util.ExitToMainConfirmation)
	fmt.Scanln(&mainMenuConfirmation)
	if mainMenuConfirmation == "y" {
		MainMenu()
	}
}
func ExitApp() {
	fmt.Println("Selamat Tinggal")
	os.Exit(0)
}
