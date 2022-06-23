package routes

import (
	"api-merca/src/repository"
	"api-merca/src/routes/handler"

	"github.com/gin-gonic/gin"
)

type Router struct {
}

func (route Router) Route() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.SetTrustedProxies(nil)

	createRoutes(r)

	r.Run(":8000")
}

// createDefaultRoutes é Responsavel pro criar as rotas padrão do CRUD, basta adicionar o controlador assinado pela interface
func createRoutes(r *gin.Engine) {
	db := repository.Basic{}

	gin.SetMode(gin.ReleaseMode)
	handler.Handler{Repo: db, Route: r}.MakeHandlers()

}
