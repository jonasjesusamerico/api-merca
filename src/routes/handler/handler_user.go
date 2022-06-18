package handler

import (
	"api-merca/src/controllers"
	"api-merca/src/middlewares"
	"api-merca/src/repository"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Repo  repository.IRepository
	Route *gin.RouterGroup
}

func (user UserHandler) RotasAutenticadas() UserHandler {

	controller := controllers.UserController{Repo: user.Repo}

	route := user.Route.Group(controller.NameGroupRoute(), middlewares.MiddleAuth())
	{
		route.GET("/", controller.FindAll)
		route.GET("/:id", controller.FindById)
		route.PATCH("/:id", controller.Update)
		route.DELETE("/:id", controller.Delete)
	}

	return user
}

func (user UserHandler) RotasNaoAutenticadas() UserHandler {
	controller := controllers.UserController{Repo: user.Repo}

	route := user.Route.Group(controller.NameGroupRoute())
	{
		route.POST("/", controller.Create)
	}

	return user
}
