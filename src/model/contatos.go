package model

import "api-merca/src/contexto"

type Contatos struct {
	Contacts []CellPhone
}

func (contato *Contatos) Adequar() {
	isCustomizavel := contexto.ContextoAutenticacao.IsCustomizavel()
	for i, element := range contato.Contacts {
		element.Validate()
		contato.Contacts[i] = element.Formatar(isCustomizavel)
	}
}
