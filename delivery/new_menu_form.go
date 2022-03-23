package delivery

import (
	"fmt"
	"warung_makan_bahari/delivery/util"
	"warung_makan_bahari/usecase"
)

func NewMenuForm(usecase usecase.MenuRegistrationUseCase) {
	var menuCode string
	var menuName string
	var menuPrice string
	var saveMenuConfirmation string
	util.CreateHeader(util.NewMenuFormHeader)
	fmt.Print(util.MenuCodeInput)
	fmt.Scanln(&menuCode)
	fmt.Print(util.MenuNameInput)
	fmt.Scanln(&menuName)
	fmt.Print(util.MenuPriceInput)
	fmt.Scanln(&menuPrice)
	fmt.Printf(util.SaveMenuConfirmation, menuCode, menuName, menuPrice)

	if saveMenuConfirmation == "y" {
		usecase.Register(menuCode, menuName, menuPrice)
	}
	MainMenu()
}
