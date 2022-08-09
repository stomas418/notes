package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stomas418/notes/api/controllers"
	"github.com/stomas418/notes/api/middleware"
)

func Route(router *gin.Engine, h *controllers.BaseHandler) {
	router.Use(middleware.SetUserStatus())
	Sign(router, h)
	Users(router, h)
	Notes(router, h)
}
