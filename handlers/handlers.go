package handlers

import (
	"github.com/gin-gonic/gin"
)

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
