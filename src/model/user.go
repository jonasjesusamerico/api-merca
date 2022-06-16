package model

import (
	"api-merca/src/auth"
	"api-merca/src/model/enum"
	"strings"
	"time"
)

// User representa um usuário utilizando a rede social
type User struct {
	ID             uint64          `json:"id,omitempty"`
	Email          string          `json:"email,omitempty"`
	Senha          string          `json:"senha,omitempty"`
	Banco          enum.BancoDados `json:"banco,omitempty"`
	IsCustomizavel bool            `json:"is_customizavel,omitempty"`
	ClienteName    string          `json:"cliente_name,omitempty"`
	CriadoEm       time.Time       `json:"criado_em,omitempty"`
}

func (usuario User) GetId() uint64 {
	return usuario.ID
}

func (usuario *User) Validate() error {

	return nil
}

func (usuario *User) Formatar() error {
	usuario.Email = strings.TrimSpace(usuario.Email)

	senhaComHash, erro := auth.Hash(usuario.Senha)
	if erro != nil {
		return erro
	}

	usuario.Senha = string(senhaComHash)

	return nil
}
