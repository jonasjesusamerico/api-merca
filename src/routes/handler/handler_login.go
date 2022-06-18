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

func (user LoginHandler) RotasAutenticadas() LoginHandler {
	return user
}

func (user LoginHandler) RotasNaoAutenticadas() LoginHandler {

	controller := controllers.LoginController{Repo: user.Repo}

	route := user.Route.Group("/" + controller.NameGroupRoute())
	{
		route.POST("/", controller.Login)
	}

	return user
}
