package model

import (
	"api-merca/src/auth"
	"api-merca/src/contexto"
	"api-merca/src/model/enum"
	"testing"
	"time"
)

func TestGetIdUsuario(t *testing.T) {
	usuario := validUsuario()

	if usuario.GetId() != 1 {
		t.Errorf("Usuario está com ID inválido")
	}
}

func TestValidateUsuario(t *testing.T) {
	contexto.CriaContextoGlobalAutenticacao()
	contexto.SetContextoAutenticacao(1, string(enum.POSTGRES_SQL), true)

	usuario := validUsuario()
	usuario.Validar()

	if usuario.GetTenantId() != 1 {
		t.Errorf("Usuario está com tenantId inválido")
	}

	if usuario.Email == " silvasauro@email.com " {
		t.Errorf("Erro ao remover caracter em branco no início e final do email do usuário")
	}

	err := auth.VerificarSenha(usuario.Senha, "!@#$%+_)(*")
	if err != nil {
		t.Errorf("Senha encriptada de forma inesperada")
	}
}

func validUsuario() Usuario {
	return Usuario{
		ID:             1,
		Email:          " silvasauro@email.com ",
		Senha:          "!@#$%+_)(*",
		IsCustomizavel: true,
		ClienteName:    "Silva Sauro TI LTDA",
		CriadoEm:       time.Now(),
		BancoDados:     string(enum.POSTGRES_SQL),
		Tenant:         Tenant{TenantId: 1},
	}
}
