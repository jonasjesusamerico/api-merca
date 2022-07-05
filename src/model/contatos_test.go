package model

import (
	"api-merca/src/contexto"
	"api-merca/src/model/enum"
	"testing"
)

func TestAdequar(t *testing.T) {
	contexto.CriaContextoGlobalAutenticacao()
	contexto.SetContextoAutenticacao(1, string(enum.POSTGRES_SQL), true)
	var telefones []Telefone

	telefone := Telefone{
		ID:   1,
		Name: "Silva Sauro",
		Num:  "0000000000000",
	}

	telefones = append(telefones, telefone)
	contatos := Contatos{Contacts: telefones}

	contatos.Adequar()

	if contatos.Contacts[0].Num != "+00 (00) 00000-0000" {
		t.Errorf("Erro ao customizar o número do telefone")
	}

	if contatos.Contacts[0].Name != "SILVA SAURO" {
		t.Errorf("Erro ao customizar o nome do telefone")
	}
}
