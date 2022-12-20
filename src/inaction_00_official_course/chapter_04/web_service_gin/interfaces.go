package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Album struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{1, "Blue Train", "John Coltrane", 56.99},
	{2, "Jeru", "Gerry Mulligan", 17.99},
	{3, "Sarah Vaughan and Clifford Brown", "Sarah Vaughan", 39.99},
}

func selectAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

/*
 * {"id": 4,"title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99} 使用 REST Client 测试!
 */
func addAlbum(c *gin.Context) {
	var album Album
	if err := c.BindJSON(&album); err != nil {
		return
	}

	albums = append(albums, album)
	c.IndentedJSON(http.StatusCreated, album)
}

/**
 * 路径参数案例
 */
func selectByIdPathParam(c *gin.Context) {

	param := c.Param("id")
	fmt.Sprintf("selectByIdPathParam method invoke, param is :", param)

	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		fmt.Errorf("selectByIdPathParam method invoke, parse int error, %p, %v", param, err)
		return
	}

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}

/**
 * 查询参数案例
 */
func selectByIdQueryParam(c *gin.Context) {

	param := c.Request.URL.Query().Get("id")
	fmt.Sprintf("selectByIdQueryParam method invoke, param is :", param)

	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		fmt.Errorf("selectByIdQueryParam method invoke, parse int error, %p, %v", param, err)
		return
	}

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}
