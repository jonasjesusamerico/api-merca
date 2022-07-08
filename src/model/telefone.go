package model

import (
	"errors"
	"strings"
)

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

func (t Telefone) Formatar(customizar bool) (telefone Telefone, err error) {

	numero := strings.Replace(t.Num, "[^\\d.]", "", -1)
	if len(numero) != 13 {
		errorMessage := "O telefone de: " + t.Name + " com o número: " + t.Num + ", não está de acordo com o padrão do sistema. +00 (00) 00000-0000. Ou com pelo menos 13 digitos válidos"
		err = errors.New(errorMessage)
		return
	}

	if !customizar {
		telefone = t
		return
	}

	pais := numero[0:2]
	ddd := numero[2:4]
	part1 := numero[4:9]
	part2 := numero[9:]

	t.Num = "+" + pais + " (" + ddd + ") " + part1 + "-" + part2
	t.Name = strings.ToUpper(t.Name)
	telefone = t
	return
}
