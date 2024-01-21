package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type Mela struct {
	Kind      string `dynamo:"kind,hash"`
	ID        string `dynamo:"id,range"`
	Point     int    `dynamo:"point" localIndex:"point-index,range"`
	SpendBy   string `dynamo:"user_id"`
	SpendTo   string `dynamo:"content_id"`
	CreatedAt string `dynamo:"created_at" localIndex:"created_at-index,range"`
}

type DynamoClient struct {
	Session *session.Session
	db      *dynamo.DB
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







