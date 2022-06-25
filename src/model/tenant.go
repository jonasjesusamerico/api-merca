package model

import (
	"api-merca/src/contexto"
)

type Tenant struct {
	TenantId   uint64 `json:"tenant_id,omitempty"`
	BancoDados string `json:"banco_dados,omitempty"`
}

func (t *Tenant) SetTenant() {
	t.TenantId = contexto.ContextoAutenticacao.GetTenantId()
	if t.BancoDados == "" {
		t.BancoDados = contexto.ContextoAutenticacao.GetBancoDados()
	}
}
