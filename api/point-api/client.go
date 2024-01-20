package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type dynamoClient struct {
	Session *session.Session
	db      *dynamo.DB
}

func (client *dynamoClient) getDB() *dynamo.DB {
	if client.db == nil {
		client.db = dynamo.New(client.Session)
	}
	return client.db
}

// Get user
func (client *dynamoClient) getUser(id string) (User, error) {
	var user User
	table := client.getDB().Table("user")
	err := table.Get("id", id).One(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Get users
func (client *dynamoClient) getUsers() ([]User, error) {
	var users []User
	table := client.getDB().Table("user")
	err := table.Scan().All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Put user
func (client *dynamoClient) putUser(user User) error {
	table := client.getDB().Table("user")
	return table.Put(user).Run()
}

// Get content
func (client *dynamoClient) getContent(id string) (Content, error) {
	var content Content
	table := client.getDB().Table("content")
	err := table.Get("id", id).One(&content)
	if err != nil {
		return content, err
	}
	return content, nil
}

// Get contents
func (client *dynamoClient) getContents() ([]Content, error) {
	var contents []Content
	table := client.getDB().Table("content")
	err := table.Scan().All(&contents)
	if err != nil {
		return nil, err
	}
	return contents, nil
}


// Put content
func (client *dynamoClient) putContent(content Content) error {
	table := client.getDB().Table("content")
	return table.Put(content).Run()
}

// Get speinding
func (client *dynamoClient) getSpending(content_id, user_id string) (Spending, error) {
	var spending Spending
	table := client.getDB().Table("spending")
	err := table.Get("content_id", content_id).Range("user_id", dynamo.Equal, user_id).One(&spending)
	if err != nil {
		return spending, err
	}
	return spending, nil
}

// Get spendings
func (client *dynamoClient) getSpendings(content_id string) ([]Spending, error) {
	var spendings []Spending
	table := client.getDB().Table("spending")
	err := table.Scan().Filter("content_id", dynamo.Equal, content_id).All(&spendings)
	if err != nil {
		return nil, err
	}
	return spendings, nil
}

// Put spending
func (client *dynamoClient) putSpending(spending Spending) error {
	table := client.getDB().Table("spending")
	return table.Put(spending).Run()
}