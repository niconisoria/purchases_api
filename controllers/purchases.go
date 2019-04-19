package controllers

import (
	"fmt"
	"net/http"
	"workshop/models"
	"workshop/services"
	"workshop/tools"

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
	c.JSON(http.StatusOK, services.GetAllPurchases())
}

func UpdatePurchase(c *gin.Context) {
	id := c.Param("id")
	purchase := models.Purchase{}

	if err := c.BindJSON(&purchase); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid data type or format")
		return
	}
	newPurchase, err := services.UpdatePurchase(id, purchase)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, newPurchase)
}

func ReadPurchases(c *gin.Context) {
	id := c.Param("id")

	if isValid := tools.ValidateString(id); !isValid {
		c.JSON(http.StatusBadRequest, "invalid params")
		return
	}

	if purchase, err := services.GetPurchaseByID(id); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(err))
		return
	} else {
		c.JSON(http.StatusOK, purchase)
		return
	}
}

func DeletePurchase(c *gin.Context) {
	id := c.Param("id")

	if isValid := tools.ValidateString(id); !isValid {
		c.JSON(http.StatusBadRequest, "invalid params")
		return
	}
	c.JSON(http.StatusAccepted, fmt.Sprint(services.DeletePurchase(id)))
	return
}
