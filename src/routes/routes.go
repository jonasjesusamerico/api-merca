package routes

import (
	"api-merca/src/routes/handler"

	"github.com/gin-gonic/gin"
)

type Router struct {
}

func (route Router) Route() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.SetTrustedProxies(nil)

	gin.SetMode(gin.ReleaseMode)
	handler.Handler{Route: r}.MakeHandlers()

	r.Run(":8000")
}
