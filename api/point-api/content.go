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

type ContentRepository interface {
	GetContent(id string) (Content, error)
	GetContents() ([]Content, error)
	GetContentsRanking() ([]Content, error)
	CreateContent(id string) (Content, error)
	UpdateContent(content Content) (Content, error)
}

type ContentRepositoryImpl struct {
	client *DynamoClient
}

func NewContentRepository(client *DynamoClient) (ContentRepository, error) {
	return &ContentRepositoryImpl{
		client: client,
	}, nil
}

var ContentNotFoundError = errors.New("Content not found")

// Get content
func (repo *ContentRepositoryImpl) GetContent(id string) (Content, error) {
	var content Content
	err := repo.client.db.Table("point").Get("kind", "content").Range("id", dynamo.Equal, id).One(&content)
	if err == dynamo.ErrNotFound {
		return Content{}, ContentNotFoundError
	} else if err != nil {
		return Content{}, err
	}
	return content, nil
}

// Get contents
func (repo *ContentRepositoryImpl) GetContents() ([]Content, error) {
	var contents []Content
	err := repo.client.db.Table("point").Get("kind", "content").All(&contents)
	if err != nil {
		return []Content{}, err
	}
	return contents, nil
}

// Get contents ranking
func (repo *ContentRepositoryImpl) GetContentsRanking() ([]Content, error) {
	var contents []Content
	err := repo.client.db.Table("point").Get("kind", "content").Index("point-index").Order(dynamo.Descending).All(&contents)
	if err != nil {
		return []Content{}, err
	}
	return contents, nil
}

// Create content
func (repo *ContentRepositoryImpl) CreateContent(id string) (Content, error) {
	content := Content{
		Kind:  "content",
		Id:    id,
		Point: 0,
	}
	err := repo.client.db.Table("point").Put(content).Run()
	if err != nil {
		return Content{}, err
	}
	return content, nil
}

// Update content
func (repo *ContentRepositoryImpl) UpdateContent(content Content) (Content, error) {
	err := repo.client.db.Table("point").Put(content).Run()
	if err != nil {
		return Content{}, err
	}
	return content, nil
}
