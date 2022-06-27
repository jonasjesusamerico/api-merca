package handler

import (
	"api-merca/src/controllers"
	"api-merca/src/repository"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	Repo  repository.IRepository
	Route *gin.RouterGroup
}

func (l LoginHandler) RotasAutenticadas() LoginHandler {
	return l
}

func (l LoginHandler) RotasNaoAutenticadas() LoginHandler {

	controller := controllers.LoginController{Repo: l.Repo}

	route := l.Route.Group("/" + controller.NameGroupRoute())
	{
		route.POST("/", controller.Login)
	}

	return l
}
