package middlewares

import (
	"api-merca/src/auth"
	"fmt"
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

// Autenticar verifica se o usuário fazendo a requisição está autenticado
// func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if erro := auth.ValidarToken(r); erro != nil {
// 			// respostas.Erro(w, http.StatusUnauthorized, erro)
// 			return
// 		}
// 		proximaFuncao(w, r)
// 	}
// }

func MiddleAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if erro := auth.ValidarToken(ctx); erro != nil {
			// respostas.Erro(w, http.StatusUnauthorized, erro)
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		fmt.Println("OIII")
	}
}
