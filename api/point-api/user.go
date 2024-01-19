package main

import (
	"errors"

	"github.com/guregu/dynamo"
)

type User struct {
	Id    string `dynamo:"id,hash"`
	Point int    `dynamo:"point"`
}

func createUser(user_id string) (User, error) {
	db := getDB()
	table := db.Table("user")
	user := User{
		Id:    user_id,
		Point: 10000,
	}
	err := table.Put(user).Run()
	if err != nil {
		return user, err
	}
	return user, nil
}

func getUser(user_id string) (User, error) {
	db := getDB()
	table := db.Table("user")
	var user User
	err := table.Get("id", user_id).One(&user)
	if err == dynamo.ErrNotFound {
		return createUser(user_id)
	} else if err != nil {
		return user, err
	}
	return user, nil
}

func (user *User) spend(content_id string, point int) error {
	var err error
	err = user.pay(point)
	if err != nil {
		return err
	}
	content, err := getContent(content_id)
	if err != nil {
		return err
	}
	err = content.incrementTotal(point)
	if err != nil {
		return err
	}
	spending, err := getSpending(content_id, user.Id)
	if err != nil {
		return err
	}
	err = spending.incrementPoint(point)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) pay(point int) error {
	var err error
	if err != nil {
		return errors.New("user not found")
	}
	if user.Point < point {
		return errors.New("not enough point")
	}
	user.Point -= point
	err = user.update(user.Point - point)
	if err != nil {
		return errors.New("update user failed")
	}
	return nil
}

func (user *User) update(point int) error {
	db := getDB()
	table := db.Table("user")
	err := table.Put(user).Run()
	if err != nil {
		return err
	}
	return nil
}
