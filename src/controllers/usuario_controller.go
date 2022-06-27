package controllers

import (
	"api-merca/src/model"
	"api-merca/src/repository"
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
	// uc.Repo.Find(&usuarios)
	uc.Repo.FindAll(&usuarios, "")
	c.JSON(http.StatusOK, usuarios)
}

func (uc UsuarioController) FindById(c *gin.Context) {
	var usuario model.Usuario
	id, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 64)

	repository.Basic{}.FindById(&usuario, id)

	if usuario.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"not_found": "Usuario not found",
		})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func (uc UsuarioController) Create(c *gin.Context) {
	var usuario model.Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	uc.Repo.Save(&usuario)
	c.JSON(http.StatusOK, usuario)
}

func (uc UsuarioController) Update(c *gin.Context) {
	var usuario model.Usuario
	id := c.Params.ByName("id")

	uc.Repo.FindById(&usuario, id)

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func (uc UsuarioController) Delete(c *gin.Context) {
	var usuario model.Usuario
	id := c.Params.ByName("id")
	repository.Basic{}.Delete(&usuario, id)
	c.JSON(http.StatusOK, gin.H{"data": "Usuario deletado com sucesso"})
}

func (uc UsuarioController) RotaCustomizada(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Message": "Oi, eu sou uma rota customizada!"})
}
