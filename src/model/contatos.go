package model

import "api-merca/src/contexto"

type Contatos struct {
	Contacts []Telefone
}

func (contato *Contatos) Adequar() (erro error) {
	isCustomizavel := contexto.ContextoAutenticacao.IsCustomizavel()
	for i, telefone := range contato.Contacts {
		telefone.Validate()
		tel, err := telefone.Formatar(isCustomizavel)
		if err != nil {
			erro = err
			return
		}
		contato.Contacts[i] = tel
	}
	return
}
