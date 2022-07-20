package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stomas418/notes/api/controllers"
	"github.com/stomas418/notes/api/database"
	"github.com/stomas418/notes/api/middleware"
	"github.com/stomas418/notes/api/routes"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	h := controllers.NewBaseHandler(db)

	router := gin.Default()
	router.Use(middleware.SetUserStatus())
	routes.Route(router, h)
	router.Run("localhost:8080")
}
