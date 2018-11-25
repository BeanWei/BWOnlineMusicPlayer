package main

import (
	"github.com/BeanWei/BWOnlineMusicPlayer/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.LoadHTMLGlob("templates/*")

	router.GET("/player", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// api group v1
	v1 := router.Group("/api/v1")
	{
		v1.POST("/music", controllers.MusicApiHandler)
	}

	router.Run(":8080")
}
