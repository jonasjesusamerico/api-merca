package model

import (
	"api-merca/src/auth"
	"strings"
	"time"

	"gorm.io/gorm"
)

// User representa um usuário utilizando a rede social
type User struct {
	ID             uint64    `json:"id,omitempty"`
	Email          string    `json:"email,omitempty"`
	Senha          string    `json:"senha,omitempty"`
	IsCustomizavel bool      `json:"is_customizavel,omitempty"`
	ClienteName    string    `json:"cliente_name,omitempty"`
	CriadoEm       time.Time `json:"criado_em,omitempty"`
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
