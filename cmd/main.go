package main

import (
	"fastfood-api/config"
	"fastfood-api/internal/handlers"
	"fastfood-api/internal/repository"
	"fastfood-api/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()
	db := database.Init()
	menuRepo := repository.NewMenuRepository(db)
	menuHandler := handlers.NewMenuHandler(menuRepo)

	r := gin.Default()
	r.GET("/menu", menuHandler.GetAllItems)
	r.GET("/menu/:id", menuHandler.GetItem)
	r.POST("/menu", menuHandler.CreateItem)
	r.PUT("/menu/:id", menuHandler.UpdateItem)
	r.DELETE("/menu/:id", menuHandler.DeleteItem)

	r.Run(":8080")
}