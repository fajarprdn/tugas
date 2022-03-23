package delivery

import (
	"fmt"
	"warung_makan_bahari/delivery/util"
	"warung_makan_bahari/usecase"
)

func ListProductForm(useCase usecase.SearchMenuUseCase) {
	util.CreateHeader(util.ListMenuFormHeader)
	util.CreateHeaderTable()
	for idx, product := range useCase.Search("") {
		fmt.Printf(util.MenuListTableFormat, idx+1, product.GetMenuCode(), product.GetMenuName(), product.GetMenuPrice())
	}
	ExitToMain()
}
