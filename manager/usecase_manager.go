package manager

import "warung_makan_bahari/usecase"

type UseCaseManager interface {
	MenuRegistrationUseCase() usecase.MenuRegistrationUseCase
	SearchMenuUseCase() usecase.SearchMenuUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) MenuRegistrationUseCase() usecase.MenuRegistrationUseCase {
	return usecase.NewMenuRegistrationUseCase(u.repo.MenuRepo())
}

func (u *useCaseManager) SearchMenuUseCase() usecase.SearchMenuUseCase {
	return usecase.NewSearchMenuUseCase(u.repo.MenuRepo())
}

func NewUseCaseManager(manager RepoManager) UseCaseManager {
	return &useCaseManager{repo: manager}
}
