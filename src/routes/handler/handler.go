package handler

import (
	"api-merca/src/middlewares"
	"api-merca/src/repository"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Repo  repository.IRepository
	Route *gin.Engine
}

func (h Handler) MakeHandlers() {

	main := h.Route.Group("/")
	{
		LoginHandler{Repo: h.Repo, Route: main}.RotasAutenticadas().RotasNaoAutenticadas()
	}

	api := main.Group("api")

	v1 := api.Group("v1", middlewares.MiddleRecriaContexto())
	{
		UserHandler{Repo: h.Repo, Route: v1}.RotasAutenticadas().RotasNaoAutenticadas()
		CellPhoneHandler{Repo: h.Repo, Route: v1}.RotasAutenticadas().RotasNaoAutenticadas()
	}
}
