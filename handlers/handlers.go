package handlers

import (
	"github.com/gin-gonic/gin"
)

var apiKeyGlobal = "1234abcd"

// return true if is valid
func validateApiKey(c *gin.Context) bool {
	xApiKeySlice, existApiKey := c.Request.Header["X-Api-Key"]
	if !existApiKey {
		c.JSON(401, gin.H{
			"message": "X-Api-Key is required",
		})
		return false
	}
	// verificar si el slice xApiKeySlice tiene al menos un valor
	if len(xApiKeySlice) == 0 {
		c.JSON(401, gin.H{
			"message": "X-Api-Key missing value",
		})
		return false
	}

	xApiKey := xApiKeySlice[0]
	if xApiKey != apiKeyGlobal {
		c.JSON(401, gin.H{
			"message": "X-Api-Key not match",
		})
		return false
	}
	return true
}

func HandlerSpanish(c *gin.Context) {
	isValid := validateApiKey(c)
	if !isValid {
		return
	}
	c.JSON(200, gin.H{
		"message": "hola",
	})
}

func HandlerEnglish(c *gin.Context) {
	isValid := validateApiKey(c)
	if !isValid {
		return
	}
	c.JSON(200, gin.H{
		"message": "hello",
	})
}

func HandlerPortuguese(c *gin.Context) {
	isValid := validateApiKey(c)
	if !isValid {
		return
	}
	c.JSON(200, gin.H{
		"message": "ola",
	})
}
