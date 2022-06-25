package model

import (
	"api-merca/src/auth"
	"strings"
	"time"

	"gorm.io/gorm"
)

// User representa um usuário utilizando a rede social
type User struct {
	ID             uint64    `gorm:"primarykey column:id" json:"id,omitempty"`
	Email          string    `gorm:"column:email" json:"email,omitempty"`
	Senha          string    `gorm:"column:senha" json:"-"`
	IsCustomizavel bool      `gorm:"column:is_customizavel" json:"is_customizavel,omitempty"`
	ClienteName    string    `gorm:"column:cliente_name" json:"cliente_name,omitempty"`
	CriadoEm       time.Time `gorm:"column:criado_em" json:"criado_em,omitempty"`
	BancoDados     string    `gorm:"column:banco_dados" json:"-"`
	Tenant
}

func (usuario User) GetId() uint64 {
	return usuario.ID
}

func (usuario *User) Validate() error {
	usuario.SetTenant()
	usuario.Formatar()

	return nil
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	err = tx.Model(u).Update("TenantId", u.ID).Error
	return
}

func (usuario *User) Formatar() (erro error) {
	usuario.Email = strings.TrimSpace(usuario.Email)

	senhaComHash, erro := auth.Hash(usuario.Senha)
	if erro != nil {
		return
	}

	usuario.Senha = string(senhaComHash)

	return
}
