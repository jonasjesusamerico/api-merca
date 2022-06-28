package controllers

import (
	"api-merca/src/auth"
	"api-merca/src/controllers/resposta"
	"api-merca/src/model"
	"api-merca/src/model/enum"
	"api-merca/src/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	Repo repository.IRepository
}

func (lc LoginController) NameGroupRoute() string {
	return "/login"
}

func (lc LoginController) Login(c *gin.Context) {
	var usuario model.Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		resposta.Erro(c, http.StatusBadRequest, err)
		return
	}

	var usuarioSalvoNoBanco model.Usuario
	erro := lc.Repo.FindFirst(&usuarioSalvoNoBanco, "email = ?", usuario.Email)
	if erro != nil {
		resposta.Erro(c, http.StatusInternalServerError, erro)
		return
	}

	if erro = auth.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); erro != nil {
		resposta.Erro(c, http.StatusUnauthorized, erro)
		return
	}

	token, erro := auth.CriarToken(usuarioSalvoNoBanco.ID, usuarioSalvoNoBanco.IsCustomizavel, enum.BancoDados(usuarioSalvoNoBanco.BancoDados))
	if erro != nil {
		resposta.Erro(c, http.StatusInternalServerError, erro)
		return
	}

	usuarioID := strconv.FormatUint(usuarioSalvoNoBanco.ID, 10)

	c.JSON(http.StatusOK, model.DadosAutenticacao{ID: usuarioID, Token: token})
}
