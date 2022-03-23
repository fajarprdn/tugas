package delivery

import (
	"fmt"
	"strings"
	"warung_makan_bahari/delivery/util"
	"warung_makan_bahari/usecase"
)

func SearchMenuForm(useCase usecase.SearchMenuUseCase) {
	var menuCode string
	//var menuName string
	util.CreateHeader(util.SearchMenuByFormHeader)
	fmt.Print(util.MenuCodeInput)
	fmt.Scanln(&menuCode)
	menu := useCase.Search(menuCode)
	if menu != nil {
		util.CreateHeaderTable()
		p := menu[0]
		fmt.Println(strings.Repeat("=", 50))
		fmt.Printf(util.MenuListTableFormat, "||", 1, "||", p.GetMenuCode(), "||", p.GetMenuName(), "||", p.GetMenuPrice(), "||")
		fmt.Println(strings.Repeat("=", 50))

	} else {
		fmt.Println("Menu Tidak Tersedia")
	}
	ExitToMain()
}
