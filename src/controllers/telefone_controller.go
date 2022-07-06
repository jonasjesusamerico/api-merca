package controllers

import (
	"api-merca/src/controllers/resposta"
	"api-merca/src/model"
	"api-merca/src/repository"
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

	err := cp.Repo.FindAll(&cellPhones, "")
	if err != nil {
		resposta.Erro(c, http.StatusInternalServerError, err)
		return
	}

	if len(cellPhones) == 0 {
		resposta.JSON(c, http.StatusNoContent, cellPhones)
		return
	}

	resposta.JSON(c, http.StatusOK, cellPhones)
}

func (cp TelefoneController) FindById(c *gin.Context) {
	var cellPhone model.Telefone
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		resposta.Erro(c, http.StatusInternalServerError, err)
		return
	}

	err = cp.Repo.FindById(&cellPhone, id)
	if err != nil {
		resposta.Erro(c, http.StatusInternalServerError, err)
		return
	}

	if cellPhone.ID == 0 {
		resposta.NotFound(c, "Telefone")
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

	err := cp.Repo.FindById(&cellPhone, id)
	if err != nil {
		resposta.Erro(c, http.StatusInternalServerError, err)
		return
	}

	if err := c.ShouldBindJSON(&cellPhone); err != nil {
		resposta.Erro(c, http.StatusBadRequest, err)
		return
	}

	resposta.JSON(c, http.StatusOK, cellPhone)
}

func (cp TelefoneController) Delete(c *gin.Context) {
	var cellPhone model.Telefone
	id := c.Params.ByName("id")
	err := cp.Repo.Delete(&cellPhone, id)
	if err != nil {
		resposta.Erro(c, http.StatusInternalServerError, err)
		return
	}
	resposta.OkMessage(c, "Telefone deletado com sucesso")
}

func (cp TelefoneController) CreateContatos(c *gin.Context) {

	var contatos model.Contatos
	if err := c.ShouldBindJSON(&contatos); err != nil {
		resposta.Erro(c, http.StatusBadRequest, err)
		return
	}

	if err := contatos.Adequar(); err != nil {
		resposta.Erro(c, http.StatusBadRequest, err)
		return
	}
	err := cp.Repo.SaveAll(contatos.Contacts)
	if err != nil {
		resposta.Erro(c, http.StatusInternalServerError, err)
		return
	}

	resposta.JSON(c, http.StatusOK, contatos.Contacts)
}
