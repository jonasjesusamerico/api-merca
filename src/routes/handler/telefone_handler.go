package handler

import (
	"api-merca/src/controllers"
	"api-merca/src/middlewares"
	"api-merca/src/repository"

	"github.com/gin-gonic/gin"
)

type TelefoneHandler struct {
	Repo  repository.IRepository
	Route *gin.RouterGroup
}

func (th TelefoneHandler) New(repo repository.IRepository, rota *gin.RouterGroup) IHandler {
	th.Repo = repo
	th.Route = rota
	return &th
}

func (th TelefoneHandler) RotasAutenticadas() IHandler {

	controller := controllers.TelefoneController{Repo: th.Repo}

	route := th.Route.Group(controller.NameGroupRoute(), middlewares.MiddleAuth())
	{
		route.GET("/", controller.FindAll)
		route.GET("/:id", controller.FindById)
		route.POST("/", controller.Create)
		route.PATCH("/:id", controller.Update)
		route.DELETE("/:id", controller.Delete)
		route.POST("/contatos", controller.CreateContatos)
	}

	return th
}

func (t TelefoneHandler) RotasNaoAutenticadas() IHandler {

	return t
}
