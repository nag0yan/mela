package main

import "testing"

func TestUpdate(t *testing.T) {
	content := Content{
		Id:    "test",
		Total: 100,
	}
	content.Total += 100
	content.update()
	if content.Total != 200 {
		t.Errorf("update failed")
	}
}