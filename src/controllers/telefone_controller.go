package controllers

import (
	"api-merca/src/model"
	"api-merca/src/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CellPhoneController struct {
	Repo repository.IRepository
}

func (cp CellPhoneController) NameGroupRoute() string {
	return "/cell-phones"
}

func (cp CellPhoneController) FindAll(c *gin.Context) {
	var cellPhones []model.CellPhone
	// cp.Repo.Find(&cellPhones)
	cp.Repo.FindAll(&cellPhones, "")
	c.JSON(http.StatusOK, cellPhones)
}

func (cp CellPhoneController) FindById(c *gin.Context) {
	var cellPhone model.CellPhone
	id, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 64)

	repository.Basic{}.FindById(&cellPhone, id)

	if cellPhone.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"not_found": "CellPhone not found",
		})
		return
	}

	c.JSON(http.StatusOK, cellPhone)
}

func (cp CellPhoneController) Create(c *gin.Context) {
	var cellPhone model.CellPhone

	if err := c.ShouldBindJSON(&cellPhone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	cp.Repo.Save(&cellPhone)
	c.JSON(http.StatusOK, cellPhone)
}

func (cp CellPhoneController) Update(c *gin.Context) {
	var cellPhone model.CellPhone
	id := c.Params.ByName("id")

	cp.Repo.FindById(&cellPhone, id)

	if err := c.ShouldBindJSON(&cellPhone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, cellPhone)
}

func (cp CellPhoneController) Delete(c *gin.Context) {
	var cellPhone model.CellPhone
	id := c.Params.ByName("id")
	repository.Basic{}.Delete(&cellPhone, id)
	c.JSON(http.StatusOK, gin.H{"data": "CellPhone deletado com sucesso"})
}

func (cp CellPhoneController) CreateContatos(c *gin.Context) {

	var contatos model.Contatos
	if err := c.ShouldBindJSON(&contatos); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	contatos.Adequar()
	cp.Repo.SaveAll(contatos.Contacts)

	c.JSON(http.StatusOK, contatos.Contacts)
}
