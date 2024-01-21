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

type UserRepository interface {
	GetUser(id string) (User, error)
	GetUsers() ([]User, error)
	CreateUser(id string) (User, error)
	UpdateUser(user User) (User, error)
}
type UserRepositoryImpl struct {
	client *DynamoClient
}

func NewUserRepository(client *DynamoClient) (UserRepository, error) {
	return &UserRepositoryImpl{
		client: client,
	}, nil
}

var UserNotFoundError = errors.New("User not found")

// Get user
func (repo *UserRepositoryImpl) GetUser(id string) (User, error) {
	var user User
	err := repo.client.db.Table("point").Get("kind", "user").Range("id", dynamo.Equal, id).One(&user)
	if err == dynamo.ErrNotFound {
		return User{}, UserNotFoundError
	} else if err != nil {
		return User{}, err
	}
	return user, nil
}

// Get users
func (repo *UserRepositoryImpl) GetUsers() ([]User, error) {
	var users []User
	err := repo.client.db.Table("point").Get("kind", "user").All(&users)
	if err != nil {
		return []User{}, err
	}
	return users, nil
}

// Create	user
func (repo *UserRepositoryImpl) CreateUser(id string) (User, error) {
	user := User{
		Kind:  "user",
		Id:    id,
		Point: 0,
	}
	err := repo.client.db.Table("point").Put(user).Run()
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// Update user
func (repo *UserRepositoryImpl) UpdateUser(user User) (User, error) {
	err := repo.client.db.Table("point").Put(user).Run()
	if err != nil {
		return User{}, err
	}
	return user, nil
}
