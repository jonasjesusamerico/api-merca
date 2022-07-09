package model

import (
	"api-merca/src/auth"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"gorm.io/gorm"
)

type Usuario struct {
	ID             uint64    `gorm:"primarykey column:id" json:"id,omitempty"`
	Email          string    `gorm:"column:email" json:"email,omitempty"`
	Senha          string    `gorm:"column:senha" json:"senha,omitempty"`
	IsCustomizavel bool      `gorm:"column:is_customizavel" json:"is_customizavel,omitempty"`
	ClienteName    string    `gorm:"column:cliente_name" json:"cliente_name,omitempty"`
	CriadoEm       time.Time `gorm:"column:criado_em" json:"criado_em,omitempty"`
	BancoDados     string    `gorm:"column:banco_dados" json:"banco_dados,omitempty"`
	Tenant
}

func (u Usuario) GetId() uint64 {
	return u.ID
}

func (u *Usuario) Validate() (err error) {
	u.SetTenant()
	u.Validar()

	if erro := checkmail.ValidateFormat(u.Email); erro != nil {
		return errors.New("o e-mail inserido é inválido")
	}

	if u.Email == "" {
		err = errors.New("o email é obrigatório e não pode estar em branco")
	}

	if u.Senha == "" {
		err = errors.New("a senha é obrigatório e não pode estar em branco")
	}

	if u.ClienteName == "" {
		err = errors.New("o cliente name é obrigatório e não pode estar em branco")
	}

	if u.BancoDados == "" {
		err = errors.New("o banco de dados é obrigatório e não pode estar em branco")
	}

	return
}

func (u *Usuario) AfterCreate(tx *gorm.DB) (err error) {
	err = tx.Model(u).Update("TenantId", u.ID).Error
	return
}

func (u *Usuario) Validar() (erro error) {
	u.Email = strings.TrimSpace(u.Email)

	senhaComHash, erro := auth.Hash(u.Senha)
	if erro != nil {
		return
	}

	u.Senha = string(senhaComHash)

	return
}

func (u Usuario) GetUsuarioRetorno() Usuario {
	return Usuario{
		ID:             u.ID,
		Email:          u.Email,
		Senha:          "",
		IsCustomizavel: u.IsCustomizavel,
		ClienteName:    u.ClienteName,
		CriadoEm:       u.CriadoEm,
		BancoDados:     "",
	}
}
