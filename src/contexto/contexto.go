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

func (ContextoGeral) GetTenantId() (tenantId uint64) {
	obj := ContextoAutenticacao.ctx.Value(Tenant{})
	if obj != nil {
		tenantId = obj.(Tenant).TenantId
	}
	return
}

func (ContextoGeral) IsCustomizavel() (isCustomizavel bool) {
	obj := ContextoAutenticacao.ctx.Value(Tenant{})
	if obj != nil {
		isCustomizavel = obj.(Tenant).IsCustomizavel
	}
	return
}

func (ContextoGeral) GetBancoDados() (bancoDados string) {
	object := ContextoAutenticacao.ctx.Value(Tenant{})
	if object != nil {
		bancoDados = object.(Tenant).BancoDados
		return
	}
	bancoDados = string(enum.POSTGRES_SQL)
	return
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
