package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nniejadlik/meli-api-categories-prices"
	"net/http"
)

func main(){
	server := gin.Default()

	server.GET("/categories/:categoryId/prices", func(c *gin.Context) {
		prices := categoriesPrices.GetPrices(c.Param("categoryId"))
		c.JSON(http.StatusOK,prices)
	})

	server.Run(":38080")
}
