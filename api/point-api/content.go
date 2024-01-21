package main

type Content struct {
	Id    string `dynamo:"id,hash"`
	Total int    `dynamo:"total" index:"total-index,hash"`
}
