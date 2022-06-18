package contexto

import (
	"api-merca/src/model/enum"
	"context"
)

var ContextoAutenticacao ContextoGeral

type Tenant struct {
	TenantId       uint64
	BancoDados     string
	IsCustomizavel bool
}

type ContextoGeral struct {
	ctx    context.Context
	cancel func()
}

func (ContextoGeral) GetTenantId() uint64 {
	obj := ContextoAutenticacao.ctx.Value(Tenant{})
	if obj == nil {
		return 0
	}
	return obj.(Tenant).TenantId
}

func (ContextoGeral) IsCustomizavel() bool {
	obj := ContextoAutenticacao.ctx.Value(Tenant{})
	if obj == nil {
		return false
	}
	return obj.(Tenant).IsCustomizavel
}

func (ContextoGeral) GetBancoDados() string {
	object := ContextoAutenticacao.ctx.Value(Tenant{})
	if object == nil {
		return string(enum.POSTGRES_SQL)
	}

	return object.(Tenant).BancoDados
}

func CriaContextoGlobalAutenticacao() {
	ContextoAutenticacao.ctx = context.Background()
	ContextoAutenticacao.ctx, ContextoAutenticacao.cancel = context.WithCancel(ContextoAutenticacao.ctx)
}

func Cancel() {
	ContextoAutenticacao.cancel()
	CriaContextoGlobalAutenticacao()
}

func SetContextoAutenticacao(tenantId uint64, bancoDados string, isCustomizavel bool) {
	ContextoAutenticacao.ctx = context.WithValue(
		ContextoAutenticacao.ctx, Tenant{},
		Tenant{
			TenantId:       tenantId,
			BancoDados:     bancoDados,
			IsCustomizavel: isCustomizavel,
		},
	)
}
