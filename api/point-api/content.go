package main

import (
	"errors"

	"github.com/guregu/dynamo"
)

type Content struct {
	Kind  string `dynamo:"kind,hash"`
	Id    string `dynamo:"id,range"`
	Point int    `dynamo:"point" localIndex:"point-index,range"`
}

var ContentNotFoundError = errors.New("Content not found")

// Create content
func (client *DynamoClient) createContent(id string) (Content, error) {
	table := client.db.Table("mela")
	content := Content{
		Kind:  "content",
		Id:    id,
		Point: 0,
	}
	err := table.Put(content).Run()
	if err != nil {
		return content, err
	}
	return content, nil
}

// Get content
func (client *DynamoClient) getContent(id string) (Content, error) {
	var content Content
	table := client.db.Table("mela")
	err := table.Get("kind", "content").Range("id", dynamo.Equal, id).One(&content)
	if err != nil {
		return content, ContentNotFoundError
	}
	return content, nil
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
