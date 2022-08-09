package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stomas418/notes/api/controllers"
	"github.com/stomas418/notes/api/middleware"
)

func Users(router *gin.Engine, h *controllers.BaseHandler) {
	user := router.Group("/:user")
	user.Use(middleware.IsLoggedIn())
	user.GET("/", h.GetUser)
	user.PUT("/", h.EditUser)
	user.DELETE("/", h.DeleteUser)
}
