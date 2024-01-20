package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/guregu/dynamo"
)

// Use struct tags much like the standard JSON library,
// you can embed anonymous structs too!

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": os.Getenv("DATABASE_ENDPOINT"),
		})
	})
	// curl -X POST -H "Content-Type: application/json" -d '{"content_id":"test","point":100,"user_id":"user1"}' localhost:8080/spend
	r.POST("/spend", func(c *gin.Context) {
		var err error
		var json Spending
		c.BindJSON(&json)
		user, err := getUser(json.UserId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			log.Print(err)
			return
		}
		err = user.spend(json.ContentId, json.Point)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			log.Print(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	// curl -X GET localhost:8080/user/user1
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		user, err := getUser(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			log.Print(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	})
	// curl -X GET localhost:8080/content/test/spendings
	r.GET("/content/:id/spendings", func(c *gin.Context) {
		id := c.Param("id")
		content, err := getContent(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			log.Print(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"spendings": content.getSpendings(),
		})
	})
	// curl -X GET localhost:8080/contents
	r.GET("/contents", func(c *gin.Context) {
		contents, err := getContents()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			log.Print(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"contents": contents,
		})
	})
	// curl -X GET localhost:8080/tables
	r.GET("/tables", func(c *gin.Context) {
		db := getDB()
		tables, err := db.ListTables().All()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			log.Print(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"tables": tables,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getDB() *dynamo.DB {
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Region: aws.String("us-east-1"), Endpoint: aws.String(os.Getenv("DATABASE_ENDPOINT"))})
	return db
}
