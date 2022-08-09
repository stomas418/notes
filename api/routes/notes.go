package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stomas418/notes/api/controllers"
	"github.com/stomas418/notes/api/middleware"
)

func Notes(router *gin.Engine, h *controllers.BaseHandler) {
	notes := router.Group("/:user/notes")
	notes.Use(middleware.IsLoggedIn())
	notes.GET("/", h.GetNotes)
	notes.POST("/", h.CreateNote)
	notes.GET("/:id", h.GetNoteById)
	notes.PUT("/:id", h.EditNote)
	notes.DELETE("/:id", h.DeleteNote)
}
