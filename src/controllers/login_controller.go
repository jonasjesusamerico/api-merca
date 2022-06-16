package controllers

import (
	"api-merca/src/auth"
	"api-merca/src/model"
	"api-merca/src/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	Repo repository.IRepository
}

func (cp LoginController) NameGroupRoute() string {
	return "/login"
}

func (cp LoginController) Login(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var usuarioSalvoNoBanco model.User
	erro := cp.Repo.FindFirst(&usuarioSalvoNoBanco, "email = ?", user.Email)
	if erro != nil {
		// respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = auth.VerificarSenha(usuarioSalvoNoBanco.Senha, user.Senha); erro != nil {
		// respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := auth.CriarToken(usuarioSalvoNoBanco.ID, usuarioSalvoNoBanco.IsCustomizavel, usuarioSalvoNoBanco.Banco)
	if erro != nil {
		// respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	usuarioID := strconv.FormatUint(usuarioSalvoNoBanco.ID, 10)

	c.JSON(http.StatusOK, model.DadosAutenticacao{ID: usuarioID, Token: token})
}
