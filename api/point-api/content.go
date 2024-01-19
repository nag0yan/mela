package main

import (
	"errors"

	"github.com/guregu/dynamo"
)

type Content struct {
	Id    string `dynamo:"id,hash"`
	Total int    `dynamo:"total"`
}

func createContent(content_id string) error {
	db := getDB()
	table := db.Table("content")
	content := Content{
		Id:    content_id,
		Total: 0,
	}
	err := table.Put(content).Run()
	if err != nil {
		return err
	}
	return nil
}

func getContent(content_id string) (Content, error) {
	db := getDB()
	table := db.Table("content")
	var content Content
	err := table.Get("id", content_id).One(&content)
	if err == dynamo.ErrNotFound {
		createContent(content_id)
		return content, nil
	} else if err != nil {
		return content, err
	}
	return content, nil
}

func (content *Content) incrementTotal(point int) error {
	content.Total += point
	return content.update()
}

func (content *Content) update() error {
	db := getDB()
	table := db.Table("content")
	err := table.Put(content).Run()
	if err != nil {
		return errors.New("update failed")
	}
	return nil
}

func (content *Content) getSpendings() []Spending {
	db := getDB()
	table := db.Table("spending")
	var spendings []Spending
	err := table.Scan().Filter("content_id = ?", content.Id).All(&spendings)
	if err != nil {
		return nil
	}
	return spendings
}
