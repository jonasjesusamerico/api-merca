package model

type Contatos struct {
	Contacts []CellPhone
}

func (contato *Contatos) Adequar() {
	for i, element := range contato.Contacts {
		element.Validate()
		contato.Contacts[i] = element.Formatar(true, true)
	}
}
