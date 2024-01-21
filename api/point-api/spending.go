package main

import (
	"time"
)

type SpendingAction struct {
	ContentId string `json:"content_id"`
	UserId    string `json:"user_id"`
	Point     int    `json:"point"`
}
type Spending struct {
	Kind       string    `dynamo:"kind,hash"`
	SpendingId string    `dynamo:"id,range"`
	ContentId  string    `dynamo:"content_id"`
	UserId     string    `dynamo:"user_id"`
	Point      int       `dynamo:"point" localIndex:"point-index,range"`
	CreatedAt  time.Time `dynamo:"created_at" localIndex:"created_at-index,range"`
}

type SpendingRepository interface {
	CreateSpending(contentId string, userId string, point int) (Spending, error)
	GetSpendings(userId string) ([]Spending, error)
	PutSpending(spending Spending) (Spending, error)
}

type SpendingRepositoryImpl struct {
	client *DynamoClient
}

func NewSpendingRepository(client *DynamoClient) (SpendingRepository, error) {
	return &SpendingRepositoryImpl{
		client: client,
	}, nil
}

// Create spending
func (repo *SpendingRepositoryImpl) CreateSpending(contentId string, userId string, point int) (Spending, error) {
	spending := Spending{
		Kind:       "spending",
		SpendingId: userId + "-" + contentId,
		ContentId:  contentId,
		UserId:     userId,
		Point:      point,
		CreatedAt:  time.Now(),
	}
	spending, err := repo.PutSpending(spending)
	if err != nil {
		return Spending{}, err
	}
	return spending, nil
}

// Get spendings
func (repo *SpendingRepositoryImpl) GetSpendings(userId string) ([]Spending, error) {
	var spendings []Spending
	err := repo.client.db.Table("point").Scan().Filter("user_id = ?", userId).All(&spendings)
	if err != nil {
		return []Spending{}, err
	}
	return spendings, nil
}

// Put spending
func (repo *SpendingRepositoryImpl) PutSpending(spending Spending) (Spending, error) {
	err := repo.client.db.Table("point").Put(spending).Run()
	if err != nil {
		return Spending{}, err
	}
	return spending, nil
}
