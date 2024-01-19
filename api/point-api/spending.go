package main

import (
	"time"

	"github.com/guregu/dynamo"
)

type Spending struct {
	ContentId string    `dynamo:"content_id,hash" json:"content_id" binding:"required"`
	Point     int       `dynamo:"point" json:"point" binding:"required"`
	UserId    string    `dynamo:"user_id,range" json:"user_id" binding:"required"`
	CreatedAt time.Time `dynamo:"created_at"`
}

func getSpending(content_id string, user_id string) (Spending, error) {
	db := getDB()
	table := db.Table("spending")
	var spending Spending
	err := table.Get("content_id", content_id).Range("user_id", dynamo.Equal, user_id).One(&spending)
	if err == dynamo.ErrNotFound {
		return createSpending(content_id, user_id)
	} else if err != nil {
		return spending, err
	}
	return spending, nil
}

func createSpending(content_id string, user_id string) (Spending, error) {
	db := getDB()
	table := db.Table("spending")
	spending := Spending{
		ContentId: content_id,
		UserId:    user_id,
		Point:     0,
	}
	err := table.Put(spending).Run()
	if err != nil {
		return spending, err
	}
	return spending, nil
}

func (spending *Spending) incrementPoint(point int) error {
	spending.Point += point
	return spending.update()
}

func (spending *Spending) update() error {
	db := getDB()
	table := db.Table("spending")
	err := table.Put(spending).Run()
	if err != nil {
		return err
	}
	return nil
}
