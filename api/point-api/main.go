package main

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// Use struct tags much like the standard JSON library,
// you can embed anonymous structs too!
type Content struct {
	Id    string `dynamo:"id"`
	Total int    `dynamo:"total"`
}

type Mela struct {
	ContentId string    `dynamo:"content_id"`
	Point     int       `dynamo:"point"`
	UserId    string    `dynamo:"user_id"`
	CreatedAt time.Time `dynamo:"created_at"`
}

func main() {
	err := addPoint("test", 100, "user1")
	if err != nil {
		panic(err)
	}
}

func getDB() *dynamo.DB {
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Region: aws.String("us-east-1")})
	return db
}

func addPoint(content_id string, point int, user_id string) error {
	err := createMela(content_id, point, user_id)
	if err != nil {
		return err
	}
	current, err := getTotal(content_id)
	if err != nil {
		return err
	}
	err = updateTotal(content_id, current+point)
	if err != nil {
		return err
	}
	return nil
}

func updateTotal(content_id string, total int) error {
	db := getDB()
	table := db.Table("content")
	content := Content{
		Id:    content_id,
		Total: total,
	}
	err := table.Put(content).Run()
	if err != nil {
		return err
	}
	return nil
}

func getTotal(content_id string) (int, error) {
	db := getDB()
	table := db.Table("content")
	var content Content
	err := table.Get("id", content_id).One(&content)
	if err == dynamo.ErrNotFound {
		createContent(content_id)
		return 0, nil
	} else if err != nil {
		return 0, err
	}
	return content.Total, nil
}

func createContent(content_id string) error {
	db := getDB()
	table := db.Table("content")
	content := Content{
		Id:    content_id,
		Total: 0,
	}
	err := table.Put(content).Run()
	if err != nil {
		return err
	}
	return nil
}

func createMela(content_id string, point int, user_id string) error {
	db := getDB()
	table := db.Table("mela")
	mela := Mela{
		ContentId: content_id,
		Point:     point,
		UserId:    user_id,
		CreatedAt: time.Now(),
	}
	err := table.Put(mela).Run()
	if err != nil {
		return err
	}
	return nil
}
