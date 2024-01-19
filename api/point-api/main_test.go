package main

import(
	"testing"
)

func TestMain(m *testing.M){
	// setup
	deleteAllTable()
	createDefaultTable()
	m.Run()
}

func deleteAllTable(){
	db := getDB()
	tables, _ := db.ListTables().All()
	for _, table := range tables{
		db.Table(table).DeleteTable().Run()
	}
}

func createDefaultTable(){
	db := getDB()
	db.CreateTable("user", User{}).Run()
	db.CreateTable("content", Content{}).Run()
	db.CreateTable("spending", Spending{}).Run()
}