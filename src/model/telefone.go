package model

import "strings"

type CellPhone struct {
	ID   uint64 `gorm:"primarykey column:id"`
	Name string `gorm:"column:name" json:"name" `
	Num  string `gorm:"column:cell_phone" json:"cellphone"`
	Tenant
}

func (c CellPhone) GetId() uint64 {
	return c.ID
}

func (c *CellPhone) Validate() error {
	c.SetTenant()
	return nil
}

func (cellPhone CellPhone) Formatar(customizar bool) CellPhone {

	if !customizar {
		return cellPhone
	}

	numero := strings.Replace(cellPhone.Num, "[^\\d.]", "", -1)
	pais := numero[0:2]
	ddd := numero[2:4]
	part1 := numero[4:9]
	part2 := numero[9:]

	cellPhone.Num = "+" + pais + " (" + ddd + ") " + part1 + "-" + part2
	cellPhone.Name = strings.ToUpper(cellPhone.Name)
	return cellPhone
}
