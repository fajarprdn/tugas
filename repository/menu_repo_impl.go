package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"warung_makan_bahari/model"
)

type menuRepoImpl struct {
	menuDb *sqlx.DB
}

func (m *menuRepoImpl) Insert(newMenu model.Menu) {
	tx := m.menuDb.MustBegin()
	_, err := tx.Exec("insert into menu(code, name, price) values($1, $2, $3)", newMenu, newMenu, newMenu)
	if err != nil {
		fmt.Println(err)
	}
	tx.Commit()
}

func (m *menuRepoImpl) GetByMenuCode(menuCode string) model.Menu {
	var selectedMenu model.Menu
	m.menuDb.Get(&selectedMenu, "select * from product where code = $1", menuCode) // Get seperti first pada silverstripe
	return selectedMenu
}

//func (m *menuRepoImpl) GetByMenuName(menuName string) model.Menu {
//	var selectedMenu model.Menu
//	for _, menu := range *m.menuDb {
//		mn := menu.GetMenuName()
//		if mn == menuName {
//			selectedMenu = menu
//		}
//	}
//	return selectedMenu
//}

func (m *menuRepoImpl) GetAll() []model.Menu {
	var dataMenu []model.Menu
	sql := `select * from menu`
	m.menuDb.Select(&dataMenu, sql) // Get dengan jumlah data yang banyak
	return dataMenu
}

func NewMenuRepo(menuDb *sqlx.DB) MenuRepo {
	menuRepo := menuRepoImpl{
		menuDb: menuDb,
	}
	return &menuRepo
}
