package main

import (
	"log"
	"testing"
)

func TestMain(m *testing.M) {
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
	err := dbClient.db.CreateTable("mela", Mela{}).Run()
	if err != nil {
		log.Print(err)
	}
}

func DeleteTables() {
	tables, _ := dbClient.db.ListTables().All()
	for _, table := range tables {
		dbClient.db.Table(table).DeleteTable().Run()
	}
}
