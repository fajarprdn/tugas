package main

import (
	"fmt"
	"warung_makan_bahari/config"
	"warung_makan_bahari/delivery"
)

func main() {
	appConfig := config.NewConfig()
	delivery.MainMenu()
	for {
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			delivery.NewMenuForm(appConfig.UseCaseManager.MenuRegistrationUseCase())
			break
		case "2":
			delivery.ListProductForm(appConfig.UseCaseManager.SearchMenuUseCase())
			break
		case "3":
			delivery.SearchMenuForm(appConfig.UseCaseManager.SearchMenuUseCase())
			break
		case "4":
			delivery.ExitApp()
		}
	}

}
