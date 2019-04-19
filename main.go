package main

import (
	"net/http"
	"workshop/controllers"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func main() {
	routes()
	// var db = db.DBPurchases{}
	// for k, val := range db.GetAll() {
	// 	fmt.Printf("Key: %v - Value: %#v \n", k, val)
	// }
	router.Run(":8080")
}

func routes() {
	router.POST("/purchases", onlyAdmin, controllers.CreatePurchase)
	router.GET("/purchases", controllers.GetPurchases)
	router.GET("/purchases/:id", controllers.ReadPurchases)
	router.PUT("/purchases/:id", onlyAdmin, controllers.UpdatePurchase)
	router.DELETE("/purchases/:id", onlyAdmin, controllers.DeletePurchase)
}

func onlyAdmin(c *gin.Context) {
	if role := c.GetHeader("role"); role != "admin" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized!")
	}
}
