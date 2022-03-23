package model

import "fmt"

type Menu struct {
	code  string
	name  string
	price string
}

func (p *Menu) ToString() string {
	return fmt.Sprintf("Product Info : %s, %s, %s", p.code, p.name, p.price)
}

func (p *Menu) SetMenuCode(code string) {
	p.code = code
}

func (p *Menu) SetMenuName(name string) {
	p.name = name
}

func (p *Menu) SetMenuPrice(price string) {
	p.price = price
}

func (p *Menu) GetMenuCode() string {
	return p.code
}

func (p *Menu) GetMenuName() string {
	return p.name
}

func (p *Menu) GetMenuPrice() string {
	return p.price
}

func NewMenu(menuCode string, menuName string, menuPrice string) Menu {
	return Menu{
		code:  menuCode,
		name:  menuName,
		price: menuPrice,
	}
}
