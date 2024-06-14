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

	r.GET("/list/:id", func(c *gin.Context) {
		idParam := c.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameter",
		})
		return
		}

		//Retrieving value
		value,  exists := members[id]
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Id not found in the available list"})
			return
		}

		//Respond with key value pair
		c.JSON(http.StatusAccepted, gin.H{
			 "Key": id, "value": value,
		})
	
	})




	r.Run()
}

