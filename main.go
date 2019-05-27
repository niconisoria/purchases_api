package main

import (
	"fmt"
	"net/http"
	"os"
	"workshop/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "niconisoria"
	password = ""
	dbname   = ""
)

var router = gin.Default()

func main() {
	routes()
	if port := os.Getenv("PORT"); port != "" {
		router.Run(":" + port)
	} else {
		router.Run(":8080")
	}
}

func routes() {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, fmt.Sprint("Welcome to Purchases API"))
	})

	router.POST("/purchases", onlyAdmin, controllers.CreatePurchase)
	router.GET("/purchases", controllers.GetPurchases)
	router.GET("/purchases/:id", controllers.ReadPurchases)
	router.PUT("/purchases/:id", onlyAdmin, controllers.UpdatePurchase)
	router.DELETE("/purchases/:id", onlyAdmin, controllers.DeletePurchase)

	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.ReadUsers)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
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
