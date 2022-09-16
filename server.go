package main

import (
	"go_donationid/config"
	"go_donationid/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	db = config.ConnectDB()
)

func main() {
	defer config.CloseDB(db)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		var user []user.User

		db.Find(&user)

		c.JSON(http.StatusOK, gin.H{
			"data":    user,
			"message": "get users",
		})
	})

	router.Run(":8080")
}
