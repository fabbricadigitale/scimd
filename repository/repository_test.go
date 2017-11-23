package repository

import (
	"testing"
)

func TestCreateRepository(t *testing.T) {
	_, err := CreateRepository("mongodb://localhost:27017/test_db?maxPoolSize=100", "test_db", "resources")

	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
