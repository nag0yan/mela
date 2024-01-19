// test user functions
package main

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	user, err := createUser("test2")
	if err != nil {
		panic(err)
	}
	if user.Point != 10000 {
		t.Errorf("createUser failed")
	}
}

func TestPay(t *testing.T) {
	user, err := createUser("test2")
	if err != nil {
		panic(err)
	}
	if user.Point != 10000 {
		t.Errorf("createUser failed")
	}
	user.pay(100)
	if user.Point != 9900 {
		t.Errorf("Pay failed")
	}
}

func TestSpend(t *testing.T) {
	user, err := createUser("test2")
	if err != nil {
		panic(err)
	}
	content, err := getContent("abc")
	current := content.Total

	if user.Point != 10000 {
		t.Errorf("createUser failed")
	}
	user.spend("abc", 100)
	if user.Point != 9900 {
		t.Errorf("payment failed")
	}
	content, err = getContent("abc")
	if content.Total != current+100 {
		t.Errorf("receive failed")
	}
}

func TestGetUser(t *testing.T) {
	_, err := getUser("user1")
	if err != nil {
		panic(err)
	}
}
