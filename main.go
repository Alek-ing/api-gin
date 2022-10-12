package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/es", HandlerSpanish)

	r.GET("/en", HandlerEnglish)

	r.GET("/pt", HandlerPortuguese)

	r.Run("0.0.0.0:8081") // listen and serve on 0.0.0.0:8080
}

func HandlerSpanish(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hola",
	})
}

func HandlerEnglish(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}

func HandlerPortuguese(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ola",
	})
}
