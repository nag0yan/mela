package main

import (
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type Point struct {
	Kind      string    `dynamo:"kind,hash"`
	Id        string    `dynamo:"id,range"`
	Point     int       `dynamo:"point" localIndex:"point-index,range"`
	CreatedAt string    `dynamo:"created_at" localIndex:"created_at-index,range"`
	UpdatedAt time.Time `dynamo:"updated_at" localIndex:"updated_at-index,range"`
}

type DynamoClient struct {
	Session *session.Session
	db      *dynamo.DB
	UserRepo     UserRepository
	ContentRepo  ContentRepository
	SpendingRepo SpendingRepository
}

func (client *DynamoClient) getDB() error {
	sess, err := session.NewSession()
	if err != nil {
		return err
	}
	client.Session = session.Must(sess, err)
	if client.db == nil {
		client.db = dynamo.New(
			client.Session,
			&aws.Config{
				Region:   aws.String(os.Getenv("AWS_REGION")),
				Endpoint: aws.String(os.Getenv("DYNAMO_ENDPOINT")),
			},
		)
	}
	return nil
}
