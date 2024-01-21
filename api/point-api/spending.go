package main

import (
	"time"

	"github.com/guregu/dynamo"
)

type Spending struct {
	ContentId string    `dynamo:"content_id,hash" json:"content_id" binding:"required"`
	Point     int       `dynamo:"point" json:"point" binding:"required"`
	UserId    string    `dynamo:"user_id,range" json:"user_id" binding:"required"`
	UpdatedAt time.Time `dynamo:"updated_at"`
}

// Get speinding
func (client *DynamoClient) getSpending(content_id, user_id string) (Spending, error) {
	var spending Spending
	table := client.db.Table("spending")
	err := table.Get("content_id", content_id).Range("user_id", dynamo.Equal, user_id).One(&spending)
	if err != nil {
		return spending, err
	}
	return spending, nil
}

// Get spendings
func (client *DynamoClient) getSpendings(content_id string) ([]Spending, error) {
	var spendings []Spending
	table := client.db.Table("spending")
	err := table.Scan().Filter("content_id", dynamo.Equal, content_id).All(&spendings)
	if err != nil {
		return nil, err
	}
	return spendings, nil
}

// Put spending
func (client *DynamoClient) putSpending(spending Spending) *dynamo.Put {
	table := client.db.Table("spending")
	return table.Put(spending)
}
