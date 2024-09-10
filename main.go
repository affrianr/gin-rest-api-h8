package main

import (
	database "github.com/affrianr/gin-rest-api-h8/config"

	"github.com/affrianr/gin-rest-api-h8/handler"
	"github.com/affrianr/gin-rest-api-h8/repository"
	"github.com/affrianr/gin-rest-api-h8/usecase"

	"github.com/gin-gonic/gin"
)


func main(){
	db := database.ConnectDatabase()

	orderRepo := repository.NewPostgresOrderRepository(db)
	orderUsecase := usecase.NewOrderUseCase(orderRepo)
	orderHandler := handler.NewOrderHandler(orderUsecase)

	itemRepo := repository.NewPostgresItemRepository(db)
	itemUsecase := usecase.NewItemUseCase(itemRepo)
	itemHandler := handler.NewItemHandler(itemUsecase)

	r := gin.Default()

	r.POST("/orders", orderHandler.CreateOrder)
	r.GET("/orders/:id", orderHandler.GetOrder)
	r.PUT("/orders/:id", orderHandler.UpdateOrder)
	r.DELETE("/orders/:id", orderHandler.DeleteOrder)
	r.GET("/orders", orderHandler.ListOrders)

	r.POST("/items", itemHandler.CreateItem)
	r.GET("/items/:id", itemHandler.GetItem)
	r.PUT("/items/:id", itemHandler.UpdateItem)
	r.DELETE("/items/:id", itemHandler.DeleteItem)
	r.GET("/items", itemHandler.ListItems)


	r.Run(":8080")
	
}