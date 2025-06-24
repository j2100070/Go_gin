package main

import (
	"Go_gin/controllers"
	"Go_gin/infra"
	"Go_gin/migrations"
	"Go_gin/repositories"
	"Go_gin/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	migrations.Migrate()

	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	r := gin.Default()
	r.POST("/items", itemController.Create)
	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindByID)
	r.POST("/items/bulk_create", itemController.BulkCreate)
	r.DELETE("/items", itemController.DeleteAll)
	r.POST("/add_users", userController.Create)

	r.Run("localhost:8080")
}
