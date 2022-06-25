package model

import (
	"api-merca/src/contexto"
)

type Tenant struct {
	TenantId uint64 `gorm:"column:tenant_id; index:idx_tenant_id" json:"-"`
	// BancoDados string `json:"banco_dados,omitempty"`
}

func (t *Tenant) SetTenant() {
	t.TenantId = contexto.ContextoAutenticacao.GetTenantId()
}
