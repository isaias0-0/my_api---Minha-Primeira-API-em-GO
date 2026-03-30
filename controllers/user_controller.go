package controllers

import (
	"my_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	user := models.User{
		ID:    1,
		Name:  "Isaias",
		Email: "isaias.ctt0@gmail.com",
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
	}

	c.JSON(http.StatusCreated, user)
}
