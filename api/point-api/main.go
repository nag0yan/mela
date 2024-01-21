package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Use struct tags much like the standard JSON library,
// you can embed anonymous structs too!

var dbClient *DynamoClient

func main() {
	dbClient = &DynamoClient{}
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
	contentRepo, err := NewContentRepository(dbClient)
	userRepo, err := NewUserRepository(dbClient)
	spendingRepo, err := NewSpendingRepository(dbClient)
	dbClient.ContentRepo = contentRepo
	dbClient.UserRepo = userRepo
	dbClient.SpendingRepo = spendingRepo
	if err != nil {
		log.Print(err)
		log.Print("Failed to create repository")
	}

	// curl -X GET localhost:8080/content/test
	// If content does not exist, create new content
	r.GET("/content/:id", func(c *gin.Context) {
		contentId := c.Param("id")
		content, err := contentRepo.GetContent(contentId)
		if err == nil {
			c.JSON(http.StatusOK, content)
			return
		}
		if err == ContentNotFoundError {
			content, err = contentRepo.CreateContent(contentId)
			if err != nil {
				log.Print(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Failed to create content",
				})
			}
			log.Printf("Create content: %v", content.Id)
			c.JSON(http.StatusOK, content)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get content",
		})
	})

	// curl -X GET localhost:8080/contents
	r.GET("/contents", func(c *gin.Context) {
		contents, err := contentRepo.GetContents()
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to get contents",
			})
		}
		c.JSON(http.StatusOK, contents)
	})

	// curl -X GET localhost:8080/contents/sorted
	r.GET("/contents/ranking", func(c *gin.Context) {
		contents, err := contentRepo.GetContentsRanking()
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}
		c.JSON(http.StatusOK, contents)
	})

	// curl -X GET localhost:8080/content/test/ranking
	// Get user raning of content sorted by total spending point
	r.GET("/content/:id/ranking", func(c *gin.Context) {

	})

	// curl -X GET localhost:8080/user/user1
	// If user does not exist, create new user
	r.GET("/user/:id", func(c *gin.Context) {
		userId := c.Param("id")
		user, err := userRepo.GetUser(userId)
		if err == UserNotFoundError {
			user, err = userRepo.CreateUser(userId)
			if err != nil {
				log.Print(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Failed to create user",
				})
			}
			log.Printf("Create user: %v", user.Id)
			c.JSON(http.StatusOK, user)
			return
		}
		if err == nil {
			c.JSON(http.StatusOK, user)
			return
		}
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get user",
		})
	})

	// curl -X POST -H "Content-Type: application/json" -d '{"content_id":"test", "user_id":"user1","point":100}' localhost:8080/spend
	r.POST("/spend", func(c *gin.Context) {
		var json SpendingAction
		if err := c.BindJSON(&json); err != nil {
			log.Print(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed to bind json",
			})
			return
		}
		content, err := contentRepo.GetContent(json.ContentId)
		if err == ContentNotFoundError {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Content not found",
			})
			return
		}
		user, err := userRepo.GetUser(json.UserId)
		if err == UserNotFoundError {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "User not found",
			})
			return
		}
		spending, err := user.Spend(content, json.Point, dbClient)
		if err == NotEnoughPointError {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Not enough point",
			})
			return
		}
		c.JSON(http.StatusOK, spending)
	})
	// curl -X GET localhost:8080/user/user1/spendings
	r.GET("/user/:id/spendings", func(c *gin.Context) {
		userId := c.Param("id")
		user, err := userRepo.GetUser(userId)
		if err == UserNotFoundError {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "User not found",
			})
			return
		}
		spendings, err := spendingRepo.GetSpendings(user.Id)
		c.JSON(http.StatusOK, spendings)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to get spendings",
			})
			return
		}
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
