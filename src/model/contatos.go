package model

import "api-merca/src/contexto"

type Contatos struct {
	Contacts []Telefone
}

func (contato *Contatos) Adequar() (err error) {
	isCustomizavel := contexto.ContextoAutenticacao.IsCustomizavel()
	for i, telefone := range contato.Contacts {
		telefone.Validate()
		tel, erro := telefone.Formatar(isCustomizavel)
		if err != nil {
			err = erro
			break
		}
		contato.Contacts[i] = tel
	}
	return
}
