package model

import (
	"api-merca/src/contexto"
	"api-merca/src/model/enum"
	"testing"
)

func TestGetId(t *testing.T) {
	telefone := validTelefone()

	if telefone.GetId() != 1 {
		t.Errorf("Numero do ID do telefone é invalido")
	}
}

func TestValidate(t *testing.T) {
	contexto.CriaContextoGlobalAutenticacao()
	contexto.SetContextoAutenticacao(2, string(enum.POSTGRES_SQL), true)
	telefone := validTelefone()

	telefone.Validate()

	if telefone.GetTenantId() != 2 {
		t.Errorf("Erro ao adicionar o tenantId no telefone")
	}
}

func TestFormatar(t *testing.T) {
	contexto.CriaContextoGlobalAutenticacao()
	contexto.SetContextoAutenticacao(1, string(enum.POSTGRES_SQL), true)

	telefone := validTelefone().Formatar(true)

	if telefone.Num != "+00 (00) 00000-0000" {
		t.Errorf("Erro ao customizar o número do telefone")
	}

	if telefone.Name != "SILVA SAURO" {
		t.Errorf("Erro ao customizar o nome do telefone")
	}
}

func validTelefone() Telefone {
	return Telefone{
		ID:     1,
		Name:   "Silva Sauro",
		Num:    "0000000000000",
		Tenant: Tenant{TenantId: 1},
	}
}
