package main

import "errors"

var NotEnoughPointError = errors.New("Not enough point")

func (user User) Spend(content Content, point int, dbClient *DynamoClient) (Spending, error) {
	if user.Point < point {
		return Spending{}, NotEnoughPointError
	}
	user.Point = user.Point - point
	user, err := dbClient.UserRepo.UpdateUser(user)
	if err != nil {
		return Spending{}, err
	}
	content.Point = content.Point + point
	content, err = dbClient.ContentRepo.UpdateContent(content)
	if err != nil {
		return Spending{}, err
	}
	spending, err := dbClient.SpendingRepo.CreateSpending(content.Id, user.Id, point)
	if err != nil {
		return Spending{}, err
	}
	return spending, nil
}
