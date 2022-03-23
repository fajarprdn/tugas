package usecase

import (
	"warung_makan_bahari/model"
	"warung_makan_bahari/repository"
)

type MenuRegistrationUseCase interface {
	Register(menuCode string, menuName string, menuPrice string) model.Menu
}

type menuRegistrationUseCase struct {
	repo repository.MenuRepo
}

func (a *menuRegistrationUseCase) Register(menuCode string, menuName string, menuPrice string) model.Menu {
	newMenu := model.NewMenu(menuCode, menuName, menuPrice)
	a.repo.Insert(newMenu)
	return newMenu
}

func NewMenuRegistrationUseCase(repo repository.MenuRepo) MenuRegistrationUseCase {
	return &menuRegistrationUseCase{
		repo: repo,
	}
}
