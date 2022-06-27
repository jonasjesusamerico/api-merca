package handler

import (
	"api-merca/src/controllers"
	"api-merca/src/middlewares"
	"api-merca/src/repository"

	"github.com/gin-gonic/gin"
)

type UsuarioHandler struct {
	Repo  repository.IRepository
	Route *gin.RouterGroup
}

func (u UsuarioHandler) RotasAutenticadas() UsuarioHandler {

	controller := controllers.UsuarioController{Repo: u.Repo}

	route := u.Route.Group(controller.NameGroupRoute(), middlewares.MiddleAuthCriaContextoDefaultDataBase())
	{
		route.GET("/", controller.FindAll)
		route.GET("/:id", controller.FindById)
		route.PATCH("/:id", controller.Update)
		route.DELETE("/:id", controller.Delete)
	}

	return u
}

func (u UsuarioHandler) RotasNaoAutenticadas() UsuarioHandler {
	controller := controllers.UsuarioController{Repo: u.Repo}

	route := u.Route.Group(controller.NameGroupRoute())
	{
		route.POST("/", controller.Create)
	}

	return u
}
