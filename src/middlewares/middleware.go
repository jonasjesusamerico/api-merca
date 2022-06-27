package middlewares

import (
	"api-merca/src/auth"
	"api-merca/src/contexto"
	"api-merca/src/model/enum"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MiddleAuth() (funcao gin.HandlerFunc) {
	funcao = func(ctx *gin.Context) {
		if erro := auth.ValidarToken(ctx); erro != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		usuarioId, _ := auth.ExtrairUsuarioID(ctx)
		bancoDados, _ := auth.ExtrairBanco(ctx)
		isCustomizavel, _ := auth.ExtrairIsCustomizavel(ctx)

		contexto.SetContextoAutenticacao(usuarioId, bancoDados, isCustomizavel)

	}
	return
}

func MiddleAuthCriaContextoDefaultDataBase() (funcao gin.HandlerFunc) {
	funcao = func(ctx *gin.Context) {
		if erro := auth.ValidarToken(ctx); erro != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		usuarioId, _ := auth.ExtrairUsuarioID(ctx)
		bancoDados := string(enum.POSTGRES_SQL)
		isCustomizavel, _ := auth.ExtrairIsCustomizavel(ctx)
		contexto.SetContextoAutenticacao(usuarioId, bancoDados, isCustomizavel)
	}
	return
}

func MiddleRecriaContexto() (funcao gin.HandlerFunc) {
	funcao = func(ctx *gin.Context) {
		ctx.Next()
		contexto.Cancel()
		contexto.CriaContextoGlobalAutenticacao()
	}
	return
}
