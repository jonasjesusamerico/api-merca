package handler

import (
	"api-merca/src/middlewares"
	"api-merca/src/repository"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Route *gin.Engine
}

//MakeHandlers é reponsavel por construir os end point, tem como uma injeção a instancia do banco que faz o envio para todos controller
//Quando necessario a troca do serviço de banco, basta que o novo service respeite a assinatura da interface, que funcionará normal
func (h Handler) MakeHandlers() {
	basicRepository := repository.Basic{}

	main := h.Route.Group("/")
	{
		LoginHandler{Repo: basicRepository, Route: main}.RotasAutenticadas().RotasNaoAutenticadas()
	}

	api := main.Group("api")

	v1 := api.Group("v1", middlewares.MiddleRecriaContexto())
	{
		UsuarioHandler{Repo: basicRepository, Route: v1}.RotasAutenticadas().RotasNaoAutenticadas()
		TelefoneHandler{Repo: basicRepository, Route: v1}.RotasAutenticadas().RotasNaoAutenticadas()
	}
}
