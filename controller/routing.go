package controller

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//Route does all the routing for the app
func Route() {
	//Get port number to listen for
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	//Set up routing
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.tmpl")
	router.Static("/static", "static")

	//TODO: Split into routing groups
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})

	//TODO: Middleware for cookies/sessions
	//User Accounts ==============================================================================================================
	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.tmpl", nil)
	})
	router.POST("/register", register)

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", nil)
	})
	router.POST("/login", login)

	//Fantasy ====================================================================================================================
	router.POST("/matchResults", playMatch)

	router.Run(":" + port)
}
