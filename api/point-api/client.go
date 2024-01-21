package main

import (
	"errors"
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

// Get user
func (client *DynamoClient) getUser(id string) (User, error) {
	var user User
	table := client.db.Table("user")
	err := table.Get("id", id).One(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Get users
func (client *DynamoClient) getUsers() ([]User, error) {
	var users []User
	table := client.db.Table("user")
	err := table.Scan().All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Put user
func (client *DynamoClient) putUser(user User) *dynamo.Put {
	table := client.db.Table("user")
	return table.Put(user)
}

// Get content
func (client *DynamoClient) getContent(id string) (Content, error) {
	var content Content
	table := client.db.Table("content")
	err := table.Get("id", id).One(&content)
	if err != nil {
		return content, err
	}
	return content, nil
}

// Create content
func (client *DynamoClient) createContent(content Content) error {
	table := client.db.Table("mela")
	if err := table.Get("kind", "content").Range("id", dynamo.Equal, content.Id).One(&Mela{}); err == nil {
		return errors.New("content already exists")
	}
	err := table.Put(content).Run()
	if err != nil {
		return err
	}
	return nil
}

// Get contents
func (client *DynamoClient) getContents() ([]Content, error) {
	var contents []Content
	table := client.db.Table("mela")
	err := table.Get("kind", "content").All(&contents)
	if err != nil {
		return nil, err
	}
	return contents, nil
}

// Get contents sorted
func (client *DynamoClient) getContentsSorted() ([]Content, error) {
	var contents []Content
	table := client.db.Table("mela")
	// get contents sorted by point with total-index
	err := table.Get("kind", "content").Index("point-index").Order(dynamo.Descending).All(&contents)
	if err != nil {
		return nil, err
	}
	return contents, nil
}

// Put content
func (client *DynamoClient) putContent(content Content) *dynamo.Put {
	table := client.db.Table("content")
	return table.Put(content)
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
