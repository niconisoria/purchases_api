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
	//var a = validateString(id)
	if id == "" || user_id == "" || role == "" {
		c.JSON(http.StatusBadRequest, "invalid params")
		return
	}

	if purchase, err := services.GetPurchaseByID(id); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(err))
		return
	} else {
		c.JSON(http.StatusAccepted, purchase)
		return
	}
}
