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

// MakeCustom tem como objetivo ser adicionado as rotas que não são padrão do crud
func (th TelefoneHandler) MakeCustom(route *gin.Engine) {
	db := repository.Basic{}

	cellPhoneController := controllers.TelefoneController{Repo: db}

	r := route.Group("/" + cellPhoneController.NameGroupRoute())
	r.POST("/contatos", cellPhoneController.CreateContatos)

}

func (th TelefoneHandler) RotasAutenticadas() TelefoneHandler {

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

func (t TelefoneHandler) RotasNaoAutenticadas() TelefoneHandler {

	return t
}
