package model

import "strings"

type Telefone struct {
	ID   uint64 `gorm:"primarykey column:id"`
	Name string `gorm:"column:name" json:"name" `
	Num  string `gorm:"column:cell_phone" json:"cellphone"`
	Tenant
}

func (t Telefone) GetId() uint64 {
	return t.ID
}

func (t *Telefone) Validate() error {
	t.SetTenant()
	return nil
}

func (t Telefone) Formatar(customizar bool) Telefone {

	if !customizar {
		return t
	}

	numero := strings.Replace(t.Num, "[^\\d.]", "", -1)
	pais := numero[0:2]
	ddd := numero[2:4]
	part1 := numero[4:9]
	part2 := numero[9:]

	t.Num = "+" + pais + " (" + ddd + ") " + part1 + "-" + part2
	t.Name = strings.ToUpper(t.Name)
	return t
}
