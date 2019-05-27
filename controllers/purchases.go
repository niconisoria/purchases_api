package controllers

import (
	"fmt"
	"net/http"
	"purchases-api/models"
	"purchases-api/services"
	"purchases-api/tools"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

func CreatePurchase(c *gin.Context) {
	purchase := models.Purchase{}
	user := models.User{}

	if err := c.ShouldBindBodyWith(&purchase, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(err))
		return
	}

	if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(err))
		return
	}
	if !purchase.IsValid() {
		c.JSON(http.StatusBadRequest, "Invalid purchase params")
		return
	}
	newPurchase, err := services.CreatePurchase(purchase, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprint(err))
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
		c.JSON(http.StatusBadRequest, fmt.Sprint("Invalid id"))
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
	c.JSON(http.StatusOK, fmt.Sprint(services.DeletePurchase(id)))
}
