package main

type User struct {
	Id    string `dynamo:"id,hash"`
	Point int    `dynamo:"point"`
}