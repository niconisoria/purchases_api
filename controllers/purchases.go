package controllers

import (
	"fmt"
	"net/http"
	"workshop/models"
	"workshop/services"

	"github.com/gin-gonic/gin"
)

func CreatePurchase(c *gin.Context) {
	purchase := models.Purchase{}

	if err := c.BindJSON(&purchase); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if !purchase.IsValid() {
		c.JSON(http.StatusBadRequest, "Invalid purchase params")
		return
	}
	newPurchase, err := services.CreatePurchase(purchase)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, newPurchase)
}

func GetPurchases(c *gin.Context) {
	c.JSON(http.StatusCreated, services.GetAllPurchases())
}

func ReadPurchases(c *gin.Context) {
	id := c.Param("id")
	user_id := c.Query("user_id")
	role := c.GetHeader("role")

	if id == "" {
		c.JSON(http.StatusBadRequest, "invalid params id")
		return
	}

	c.JSON(http.StatusCreated, fmt.Sprintf("id: %v, user_id: %v, role: %v", id, user_id, role))
}
