package repository

import "warung_makan_bahari/model"

type MenuRepo interface {
	Insert(newMenu model.Menu)
	GetByMenuCode(menuCode string) model.Menu
	//GetByMenuName(menuName string) model.Menu
	GetAll() []model.Menu
}
