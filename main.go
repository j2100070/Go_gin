package main

import (
	"Go_gin/controllers"
	"Go_gin/infra"
	"Go_gin/repositories"
	"Go_gin/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r := gin.Default()
	r.POST("/items", itemController.Create)
	r.GET("/items", itemController.FindAll)
	r.Run("localhost:8080")
	r.GET("/items/id, itemController.FindByID(uint id)) 
}
