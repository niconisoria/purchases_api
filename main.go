package main

import (
	"fmt"
	"net/http"
	"workshop/controllers"
	"workshop/db"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()
var a = validateString("")

func main() {
	routes()
	var db = db.DBPurchases{}
	for k, val := range db.GetAll() {
		fmt.Printf("Key: %v - Value: %#v \n", k, val)
	}
	router.Run(":8080")
}

func routes() {
	router.POST("/purchases", onlyAdmin, controllers.CreatePurchase)
	router.GET("/purchases", controllers.GetPurchases)
	router.GET("/purchases/:id", controllers.ReadPurchases)

	router.PUT("/purchases/:id")
	router.DELETE("/purchases/:id")
}

func onlyAdmin(c *gin.Context) {
	if role := c.GetHeader("role"); role != "admin" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized!")
	}
}
