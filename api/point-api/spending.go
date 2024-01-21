package main

import "time"

type Spending struct {
	ContentId string    `dynamo:"content_id,hash" json:"content_id" binding:"required"`
	Point     int       `dynamo:"point" json:"point" binding:"required"`
	UserId    string    `dynamo:"user_id,range" json:"user_id" binding:"required"`
	UpdatedAt time.Time `dynamo:"updated_at"`
}
