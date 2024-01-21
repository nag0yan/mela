package main

import (
	"errors"

	"github.com/guregu/dynamo"
)

type User struct {
	Kind  string `dynamo:"kind,hash"`
	Id    string `dynamo:"id,range"`
	Point int    `dynamo:"point"`
}

var UserNotFoundError = errors.New("User not found")

// Get user
func (client *DynamoClient) getUser(id string) (User, error) {
	var user User
	table := client.db.Table("mela")
	err := table.Get("kind", "user").Range("id", dynamo.Equal, id).One(&user)
	if err == dynamo.ErrNotFound {
		return user, UserNotFoundError
	}
	return user, nil
}

// Create user
func (client *DynamoClient) createUser(id string) (User, error) {
	table := client.db.Table("mela")
	user := User{
		Kind:  "user",
		Id:    id,
		Point: 0,
	}
	err := table.Put(user).Run()
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
