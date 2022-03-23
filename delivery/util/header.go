package util

import (
	"fmt"
	"strings"
)

func CreateHeader(title string) {
	fmt.Println(title)
	fmt.Println(strings.Repeat("=", 50))
}

func CreateHeaderTable() {
	fmt.Printf(MenuListTableHeaderFormat, NoLabel, MenuCodeLabel, MenuNameLabel, MenuPriceLabel)
}
