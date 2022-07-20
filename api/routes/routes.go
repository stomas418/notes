package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stomas418/notes/api/controllers"
)

func Route(router *gin.Engine, h *controllers.BaseHandler) {
	Users(router, h)
	Notes(router, h)
	Sign(router, h)
}
