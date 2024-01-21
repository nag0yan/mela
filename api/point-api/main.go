package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Use struct tags much like the standard JSON library,
// you can embed anonymous structs too!

var dbClient dynamoClient

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
	// curl -X POST -H "Content-Type: application/json" -d '{"content_id":"test","point":100,"user_id":"user1"}' localhost:8080/spend
	r.POST("/spend", func(c *gin.Context) {
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
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
