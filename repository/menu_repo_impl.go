package Repository

import "warung_makan_bahari/model"

type menuRepoImpl struct {
	menuDb *[]model.Menu
}

func (m *menuRepoImpl) Insert(newMenu model.Menu) {
	*m.menuDb = append(*m.menuDb, newMenu)
}

func (m *menuRepoImpl) GetByMenuCode(menuCode string) model.Menu {
	var selectedMenu model.Menu
	for _, menu := range *m.menuDb {
		mc := menu.GetMenuCode()
		if mc == menuCode {
			selectedMenu = menu
			break
		}
		return selectedMenu
	}
}

func (m *menuRepoImpl) GetByMenuName(menuName string) model.Menu {
	var selectedMenu model.Menu
	for _, menu := range *m.menuDb {
		mn := menu.GetMenuName()
		if mn == menuName {
			selectedMenu = menu
		}
	}
	return selectedMenu
}

func (m *menuRepoImpl) GetAll() []model.Menu {
	return *m.menuDb
}

func NewProductRepo(menuDb *[]model.Menu) MenuRepo {
	menuRepo := menuRepoImpl{
		menuDb: menuDb,
	}
	return &menuRepo
}
