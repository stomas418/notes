package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stomas418/notes/api/controllers"
	"github.com/stomas418/notes/api/middleware"
)

func Sign(router *gin.Engine, h *controllers.BaseHandler) *gin.RouterGroup {
	sign := router.Group("/")
	sign.POST("/login", middleware.IsNotLoggedIn(), h.SignIn)
	sign.POST("/register", middleware.IsNotLoggedIn(), h.SignUp)
	return sign
}
