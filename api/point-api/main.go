package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Use struct tags much like the standard JSON library,
// you can embed anonymous structs too!

var dbClient DynamoClient

func main() {
	err := dbClient.getDB()
	if err != nil {
		log.Print("Failed to get DB")
		log.Print(err)
	}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": os.Getenv("DYNAMO_ENDPOINT"),
		})
	})
	// curl -X POST localhost:8080/content/test
	r.POST("/content/:id", func(c *gin.Context) {
		contentId := c.Param("id")
		var content Content
		content.Id = contentId
		content.Kind = "content"
		content.Point = 0
		err := dbClient.createContent(content)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Created content",
		})
	})
	// curl -X GET localhost:8080/contents
	r.GET("/contents", func(c *gin.Context) {
		contents, err := dbClient.getContents()
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to get contents",
			})
		}
		c.JSON(http.StatusOK, contents)
	})
	// curl -X GET localhost:8080/contents/sorted
	r.GET("/contents/sorted", func(c *gin.Context) {
		contents, err := dbClient.getContentsSorted()
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}
		c.JSON(http.StatusOK, contents)
	})

	// curl -X GET localhost:8080/user/user1
	r.GET("/user/:id", func(c *gin.Context) {
		userId := c.Param("id")
		user, err := dbClient.getUser(userId)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to get user",
			})
		}
		c.JSON(http.StatusOK, user)
	})

	// curl -X GET localhost:8080/content/test/spendings
	r.GET("/content/:id/spendings", func(c *gin.Context) {

	})
	// curl -X POST -H "Content-Type: application/json" -d '{"content_id":"test","point":100,"user_id":"user1"}' localhost:8080/spend
	r.POST("/spend", func(c *gin.Context) {
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
