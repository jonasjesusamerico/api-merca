package handler

import (
	"api-merca/src/middlewares"
	"api-merca/src/repository"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Route *gin.Engine
}

func (h Handler) MakeHandlers() {
	basicRepository := repository.Basic{}

	main := h.Route.Group("/")
	{
		LoginHandler{Repo: basicRepository, Route: main}.RotasAutenticadas().RotasNaoAutenticadas()
	}

	api := main.Group("api")

	v1 := api.Group("v1", middlewares.MiddleRecriaContexto())
	{
		UserHandler{Repo: basicRepository, Route: v1}.RotasAutenticadas().RotasNaoAutenticadas()
		CellPhoneHandler{Repo: basicRepository, Route: v1}.RotasAutenticadas().RotasNaoAutenticadas()
	}
}
