package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	members := map[int]string{1: "Aubrey", 2: "Nathan", 3: "Justin"}

	r := gin.Default()

	r.Use(cors.Default())

	r.StaticFile("/", "./index.html")

	r.GET("/list", func(c *gin.Context) {
		c.JSON(http.StatusOK, members)
	})

	//r.POST("/lists", func(c *gin.Context) {})

	r.GET("/list/:key", func(c *gin.Context) {
		idParam := c.Param("key") // Use "key" instead of "id"

		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameter"})
			return
		}

		r.POST("/list" func(c *gin.Context) {
			}

		// Retrieving value
		value, exists := members[id]
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
			return
		}

		// Respond with key value pair
		c.JSON(http.StatusOK, gin.H{
			"Key": id, "Value": value,
		})
	})

	r.Run()
}
