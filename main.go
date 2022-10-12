package main

import (
	h "github.com/Alek-ing/api-gin/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/es", h.HandlerSpanish)

	r.GET("/en", h.HandlerEnglish)

	r.GET("/pt", h.HandlerPortuguese)

	r.Run("0.0.0.0:8081") // listen and serve on 0.0.0.0:8080
}
