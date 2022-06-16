package handler

import (
	"api-merca/src/controllers"
	"api-merca/src/middlewares"
	"api-merca/src/repository"

	"github.com/gin-gonic/gin"
)

type CellPhoneHandler struct {
	Repo  repository.IRepository
	Route *gin.RouterGroup
}

// MakeCustom tem como objetivo ser adicionado as rotas que não são padrão do crud
func (ch CellPhoneHandler) MakeCustom(route *gin.Engine) {
	db := repository.Basic{}

	cellPhoneController := controllers.CellPhoneController{Repo: db}

	r := route.Group("/" + cellPhoneController.NameGroupRoute())
	r.POST("/contatos", cellPhoneController.CreateContatos)

}

func (celular CellPhoneHandler) RotasAutenticadas() CellPhoneHandler {

	controller := controllers.CellPhoneController{Repo: celular.Repo}

	route := celular.Route.Group(controller.NameGroupRoute(), middlewares.MiddleAuth())
	{
		route.GET("/", controller.FindAll)
		route.GET("/:id", controller.FindById)
		route.POST("/", controller.Create)
		route.PATCH("/:id", controller.Update)
		route.DELETE("/:id", controller.Delete)
		route.POST("/contatos", controller.CreateContatos)
	}

	return celular
}

func (celulars CellPhoneHandler) RotasNaoAutenticadas() CellPhoneHandler {

	return celulars
}
