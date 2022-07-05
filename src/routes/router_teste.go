package routes

import "github.com/gin-gonic/gin"

type RoutesTest struct {
}

func (RoutesTest) SetupRouter() (r *gin.Engine) {
	r = gin.Default()
	return
}
