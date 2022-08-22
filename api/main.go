package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stomas418/notes/api/controllers"
	"github.com/stomas418/notes/api/database"
	"github.com/stomas418/notes/api/routes"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	h := controllers.NewBaseHandler(db)
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	router := gin.Default()
	routes.Route(router, h)
	router.Run(":" + PORT)
}
