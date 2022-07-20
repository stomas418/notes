package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stomas418/notes/api/controllers"
	"github.com/stomas418/notes/api/middleware"
)

func Users(router *gin.Engine, h *controllers.BaseHandler) *gin.RouterGroup {
	user := router.Group("/:user")
	user.GET("/", middleware.IsLoggedIn(), h.GetUser)
	user.PUT("/", middleware.IsLoggedIn(), h.EditUser)
	user.DELETE("/", middleware.IsLoggedIn(), h.DeleteUser)
	return user
}
