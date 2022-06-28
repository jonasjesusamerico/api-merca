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

type TelefoneController struct {
	Repo repository.IRepository
}

func (cp TelefoneController) NameGroupRoute() string {
	return "/telefones"
}

func (cp TelefoneController) FindAll(c *gin.Context) {
	var cellPhones []model.Telefone

	cp.Repo.FindAll(&cellPhones, "")
	resposta.JSON(c, http.StatusOK, cellPhones)
}

func (cp TelefoneController) FindById(c *gin.Context) {
	var cellPhone model.Telefone
	id, _ := strconv.ParseUint(c.Params.ByName("id"), 10, 64)

	repository.Basic{}.FindById(&cellPhone, id)

	if cellPhone.ID == 0 {
		resposta.Erro(c, http.StatusNotFound, errors.New("telefone não encontrado"))
		return
	}

	resposta.JSON(c, http.StatusOK, cellPhone)
}

func (cp TelefoneController) Create(c *gin.Context) {
	var cellPhone model.Telefone

	if err := c.ShouldBindJSON(&cellPhone); err != nil {
		resposta.Erro(c, http.StatusBadRequest, err)
		return
	}

	cp.Repo.Save(&cellPhone)
	resposta.JSON(c, http.StatusOK, cellPhone)
}

func (cp TelefoneController) Update(c *gin.Context) {
	var cellPhone model.Telefone
	id := c.Params.ByName("id")

	cp.Repo.FindById(&cellPhone, id)

	if err := c.ShouldBindJSON(&cellPhone); err != nil {
		resposta.Erro(c, http.StatusBadRequest, err)
		return
	}

	resposta.JSON(c, http.StatusOK, cellPhone)
}

func (cp TelefoneController) Delete(c *gin.Context) {
	var cellPhone model.Telefone
	id := c.Params.ByName("id")
	repository.Basic{}.Delete(&cellPhone, id)
	c.JSON(http.StatusOK, gin.H{"data": "Telefone deletado com sucesso"})
	resposta.JSON(c, http.StatusOK, gin.H{"message": "Telefone deletado com sucesso"})
}

func (cp TelefoneController) CreateContatos(c *gin.Context) {

	var contatos model.Contatos
	if err := c.ShouldBindJSON(&contatos); err != nil {
		resposta.Erro(c, http.StatusBadRequest, err)
		return
	}

	contatos.Adequar()
	cp.Repo.SaveAll(contatos.Contacts)

	resposta.JSON(c, http.StatusOK, contatos.Contacts)
}
