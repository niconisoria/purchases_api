package controllers

import (
	"fmt"
	"net/http"
	"workshop/models"
	"workshop/services"
	"workshop/tools"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	user := models.User{}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if !user.IsValid() {
		c.JSON(http.StatusBadRequest, "Invalid user params")
		return
	}
	newUser, err := services.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetAllUsers())
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	user := models.User{}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid data type or format")
		return
	}
	newUser, err := services.UpdateUser(id, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, newUser)
}

func ReadUsers(c *gin.Context) {
	id := c.Param("id")

	if isValid := tools.ValidateString(id); !isValid {
		c.JSON(http.StatusBadRequest, "invalid params")
		return
	}

	if user, err := services.GetUserByID(id); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint("Invalid id"))
		return
	} else {
		c.JSON(http.StatusOK, user)
		return
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if isValid := tools.ValidateString(id); !isValid {
		c.JSON(http.StatusBadRequest, "invalid params")
		return
	}
	c.JSON(http.StatusOK, fmt.Sprint(services.DeleteUser(id)))
}
