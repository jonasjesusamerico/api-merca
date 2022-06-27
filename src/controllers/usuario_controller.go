package controllers

import (
	"api-merca/src/model"
	"api-merca/src/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Repo repository.IRepository
}

func (ur UserController) NameGroupRoute() string {
	return "/users"
}

func (ur UserController) FindAll(c *gin.Context) {
	var users []model.User
	// ur.Repo.Find(&users)
	ur.Repo.FindAll(&users, "")
	c.JSON(http.StatusOK, users)
}

func (ur UserController) FindById(c *gin.Context) {
	var user model.User
	id, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 64)

	repository.Basic{}.FindById(&user, id)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"not_found": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (ur UserController) Create(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ur.Repo.Save(&user)
	c.JSON(http.StatusOK, user)
}

func (ur UserController) Update(c *gin.Context) {
	var user model.User
	id := c.Params.ByName("id")

	ur.Repo.FindById(&user, id)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (ur UserController) Delete(c *gin.Context) {
	var user model.User
	id := c.Params.ByName("id")
	repository.Basic{}.Delete(&user, id)
	c.JSON(http.StatusOK, gin.H{"data": "User deletado com sucesso"})
}

func (ur UserController) RotaCustomizada(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Message": "Oi, eu sou uma rota customizada!"})
}
