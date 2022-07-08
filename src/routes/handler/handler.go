package handler

import (
	"api-merca/src/middlewares"
	"api-merca/src/repository"

	"github.com/gin-gonic/gin"
)

// var route *gin.Engine

type IHandler interface {
	New(Repo repository.IRepository, Route *gin.RouterGroup) IHandler
	RotasAutenticadas() IHandler
	RotasNaoAutenticadas() IHandler
}

type Handler struct {
	Route *gin.Engine
}

//MakeHandlers é reponsavel por construir os end point, tem como uma injeção a instancia do banco que faz o envio para todos controller
//Quando necessario a troca do serviço de banco, basta que o novo service respeite a assinatura da interface, que funcionará normal
func (h Handler) MakeHandlers() {
	basicRepository := repository.Basic{}

	main := h.Route.Group("/")
	api := main.Group("api")
	v1 := api.Group("v1", middlewares.MiddleRecriaContexto())

	rotasMain := []IHandler{
		&LoginHandler{},
	}

	rotasApi := []IHandler{}

	rotasV1 := []IHandler{
		&UsuarioHandler{},
		&TelefoneHandler{},
	}

	criaRotas(rotasMain, &basicRepository, main)
	criaRotas(rotasApi, &basicRepository, api)
	criaRotas(rotasV1, &basicRepository, v1)

}

func criaRotas(rotas []IHandler, repo *repository.Basic, base *gin.RouterGroup) {
	if len(rotas) == 0 {
		return
	}

	for _, rota := range rotas {
		rota.New(repo, base).RotasAutenticadas().RotasNaoAutenticadas()
	}
}
