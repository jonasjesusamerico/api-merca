package handler

import (
	"api-merca/src/model"
	"api-merca/src/model/enum"
	"api-merca/src/repository"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/goccy/go-json"
)

type MockRepository struct {
	repository.IRepository
}

func (m MockRepository) FindFirst(receiver model.IModel, query string, args ...interface{}) (err error) {
	usuario := receiver.(*model.Usuario)

	usuario.Senha = "$2a$10$07txiYCtscofBMrKBMRjhuTRG5QIWF5cRUR0g7VL3DWyVyXeqsFM."
	usuario.Email = "varejao@varejao.com"
	usuario.ID = 1
	usuario.IsCustomizavel = false
	usuario.BancoDados = string(enum.POSTGRES_SQL)

	return
}

func TestRotasNaoAutenticadas(t *testing.T) {
	r := gin.Default()
	main := r.Group("/")
	{
		LoginHandler{Repo: MockRepository{}, Route: main}.RotasAutenticadas().RotasNaoAutenticadas()
	}

	w := httptest.NewRecorder()

	user := map[string]string{"email": "varejao@varejao.com", "senha": "!@#$%+_)(*"}
	jsonUsuario, _ := json.Marshal(user)

	req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(string(jsonUsuario)))
	r.ServeHTTP(w, req)

	dadosAutenticacao := model.DadosAutenticacao{}
	json.Unmarshal(w.Body.Bytes(), &dadosAutenticacao)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "1", dadosAutenticacao.ID)

}
