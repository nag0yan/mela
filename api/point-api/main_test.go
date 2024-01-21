package main

import (
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	dbClient = &DynamoClient{}
	err := dbClient.getDB()
	if err != nil {
		log.Print("Failed to get DB")
		log.Print(err)
	}
	// Delete tables
	DeleteTables()
	// Create tables
	CreateTables()
}

func CreateTables() {
	err := dbClient.db.CreateTable("point", Point{}).Run()
	if err != nil {
		log.Print(err)
		panic(err)
	}
}

func DeleteTables() {
	tables, _ := dbClient.db.ListTables().All()
	for _, table := range tables {
		err := dbClient.db.Table(table).DeleteTable().Run()
		if err != nil {
			log.Print(err)
			panic(err)
		}
	}
}
