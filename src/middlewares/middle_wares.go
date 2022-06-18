package middlewares

import (
	"api-merca/src/auth"
	"api-merca/src/contexto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Logger escreve informações da requisição no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

func MiddleAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if erro := auth.ValidarToken(ctx); erro != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		defer func() {
			contexto.Cancel()
		}()

		usuarioId, _ := auth.ExtrairUsuarioID(ctx)
		bancoDados, _ := auth.ExtrairBanco(ctx)
		isCustomizavel, _ := auth.ExtrairIsCustomizavel(ctx)

		contexto.SetContextoAutenticacao(usuarioId, bancoDados, isCustomizavel)

	}
}
