package resposta

import (
	"github.com/gin-gonic/gin"
)

func JSON(c *gin.Context, statusCode int, dados interface{}) {

	if dados == nil {
		return
	}

	c.JSON(statusCode, dados)
}

func Erro(c *gin.Context, statusCode int, erro error) {
	JSON(c, statusCode, struct {
		Erro string `json:"error"`
	}{
		Erro: erro.Error(),
	})

}
