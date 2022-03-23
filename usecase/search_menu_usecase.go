package usecase

import (
	"warung_makan_bahari/model"
	"warung_makan_bahari/repository"
)

type SearchMenuUseCase interface {
	Search(menuCode string) []model.Menu
}

type searchMenuUseCase struct {
	repo repository.MenuRepo
}

func (a *searchMenuUseCase) Search(menuCode string) []model.Menu {
	if len(menuCode) == 0 {
		return a.repo.GetAll()
	}
	result := a.repo.GetByMenuCode(menuCode)
	if len(result.GetMenuCode()) == 0 {
		return nil
	} else {
		return []model.Menu{result}
	}
}

func NewSearchMenuUseCase(repo repository.MenuRepo) SearchMenuUseCase {
	return &searchMenuUseCase{repo: repo}
}
