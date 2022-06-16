package model

import "strings"

type CellPhone struct {
	ID   uint64 `gorm:"primarykey"`
	Name string `json:"name"`
	Num  string `json:"cellphone"`
	Tenant
}

func (c CellPhone) New(name, num string) *CellPhone {
	return &CellPhone{Name: name, Num: num}
}

func (c CellPhone) GetId() uint64 {
	return c.ID
}

func (c CellPhone) Validate() error {
	return nil
}

func (cellPhone CellPhone) Formatar(aplicaMascara, nomeMaisculo bool) CellPhone {

	if aplicaMascara && nomeMaisculo {
		numero := strings.Replace(cellPhone.Num, "[^\\d.]", "", -1)
		pais := numero[0:2]
		ddd := numero[2:4]
		part1 := numero[4:9]
		part2 := numero[9:]

		cellPhone.Num = "+" + pais + " (" + ddd + ") " + part1 + "-" + part2
		cellPhone.Name = strings.ToUpper(cellPhone.Name)

	}
	return cellPhone
}
