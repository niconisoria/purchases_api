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
	router.POST("/purchases", controllers.CreatePurchase)
	router.GET("/purchases", checkQueryParams, controllers.GetPurchases)
	router.GET("/purchases/:id", checkQueryParams, controllers.ReadPurchases)
	router.PUT("/purchases/:id", onlyAdmin, controllers.UpdatePurchase)
	router.DELETE("/purchases/:id", onlyAdmin, controllers.DeletePurchase)

	router.POST("/users")
	router.GET("/users")
	router.GET("/users/:id")
	router.PUT("/users/:id")
	router.DELETE("/users/:id")
}

func checkQueryParams(c *gin.Context) {
	if userID := c.Query("user_id"); userID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Query params user_id required.")
	}
}

func onlyAdmin(c *gin.Context) {
	if role := c.GetHeader("role"); role != "admin" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized!")
	}
}
