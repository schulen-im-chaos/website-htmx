package api

import (
	"fmt"
	"htmx/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostItem(c *gin.Context) {
	var newItem db.Item

	if err := c.ShouldBind(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, newItem)
		return
	}

	if err := db.AddItem(&newItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newItem)
}

func PostLogin(c *gin.Context) {
	var newLogin db.Login

	if err := c.BindJSON(&newLogin); err != nil {
		c.JSON(http.StatusBadRequest, newLogin)
		return
	}

	if err := db.AddLogin(&newLogin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newLogin)
}

func DeleteItem(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	book, err := db.DeleteItem(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func DeleteLogin(c *gin.Context) {
	id := c.Param("id")

	login, err := db.DeleteLogin(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, login)
}
