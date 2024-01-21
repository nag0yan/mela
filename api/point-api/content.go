package main

type Content struct {
	Kind  string `dynamo:"kind,hash"`
	Id    string `dynamo:"id,range"`
	Point int    `dynamo:"point" localIndex:"point-index,range"`
}
