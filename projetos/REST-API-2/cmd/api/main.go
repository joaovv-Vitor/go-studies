package main

import (
	"REST-API-2/cmd/internal/database"
	"REST-API-2/cmd/internal/product"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConection, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	repository := product.NewRepository(dbConection)

	usecase := product.NewUsecase(repository)

	handlers := product.NewHandler(usecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", handlers.GetProducts)
	server.POST("/product", handlers.CreateProduct)
	server.GET("/product/:productId", handlers.GetProductsById)
	server.DELETE("/product/:productId", handlers.DeleteProductByID)
	server.PUT("/product/:productId", handlers.UpdateProduct)

	server.Run(":8000")

}
