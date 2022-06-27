package model

import (
	"api-merca/src/auth"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Usuario struct {
	ID             uint64    `gorm:"primarykey column:id" json:"id,omitempty"`
	Email          string    `gorm:"column:email" json:"email,omitempty"`
	Senha          string    `gorm:"column:senha" json:"-"`
	IsCustomizavel bool      `gorm:"column:is_customizavel" json:"is_customizavel,omitempty"`
	ClienteName    string    `gorm:"column:cliente_name" json:"cliente_name,omitempty"`
	CriadoEm       time.Time `gorm:"column:criado_em" json:"criado_em,omitempty"`
	BancoDados     string    `gorm:"column:banco_dados" json:"-"`
	Tenant
}

func (u Usuario) GetId() uint64 {
	return u.ID
}

func (u *Usuario) Validate() error {
	u.SetTenant()
	u.Formatar()

	return nil
}

func (u *Usuario) AfterCreate(tx *gorm.DB) (err error) {
	err = tx.Model(u).Update("TenantId", u.ID).Error
	return
}

func (u *Usuario) Formatar() (erro error) {
	u.Email = strings.TrimSpace(u.Email)

	senhaComHash, erro := auth.Hash(u.Senha)
	if erro != nil {
		return
	}

	u.Senha = string(senhaComHash)

	return
}
