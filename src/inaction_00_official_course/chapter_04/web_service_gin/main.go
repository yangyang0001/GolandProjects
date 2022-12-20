package main

import "github.com/gin-gonic/gin"

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/selectAll", selectAll)
	router.POST("/addAlbum", addAlbum)
	router.GET("/selectByIdPathParam/:id", selectByIdPathParam)
	router.GET("/selectByIdQueryParam", selectByIdQueryParam)
	router.Run("localhost:8080")
}
