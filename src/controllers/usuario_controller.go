package controllers

import (
	"api-merca/src/controllers/resposta"
	"api-merca/src/model"
	"api-merca/src/repository"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UsuarioController struct {
	Repo repository.IRepository
}

func (uc UsuarioController) NameGroupRoute() string {
	return "/usuarios"
}

func (uc UsuarioController) FindAll(c *gin.Context) {
	var usuarios []model.Usuario

	uc.Repo.FindAll(&usuarios, "")
	c.JSON(http.StatusOK, usuarios)
	resposta.JSON(c, http.StatusOK, usuarios)
}

func (uc UsuarioController) FindById(c *gin.Context) {
	var usuario model.Usuario
	id, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 64)

	uc.Repo.FindById(&usuario, id)

	if usuario.ID == 0 {
		resposta.JSON(c, http.StatusOK, errors.New("usuário não encontrado"))
		return
	}

	resposta.JSON(c, http.StatusOK, usuario)
}

func (uc UsuarioController) Create(c *gin.Context) {
	var usuario model.Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		resposta.JSON(c, http.StatusOK, err.Error())
		return
	}

	uc.Repo.Save(&usuario)
	resposta.JSON(c, http.StatusOK, usuario)
}

func (uc UsuarioController) Update(c *gin.Context) {
	var usuario model.Usuario
	id := c.Params.ByName("id")

	uc.Repo.FindById(&usuario, id)

	if err := c.ShouldBindJSON(&usuario); err != nil {
		resposta.JSON(c, http.StatusOK, err.Error())
		return
	}

	resposta.JSON(c, http.StatusOK, usuario)
}

func (uc UsuarioController) Delete(c *gin.Context) {
	var usuario model.Usuario
	id := c.Params.ByName("id")
	uc.Repo.Delete(&usuario, id)
	resposta.JSON(c, http.StatusOK, gin.H{"message": "Usuario deletado com sucesso"})
}

func (uc UsuarioController) RotaCustomizada(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Oi, eu sou uma rota customizada!"})
}
