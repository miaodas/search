package hiendy

import (
	"testing"
)

func TestGetConnection(t *testing.T) {
	connection := GetConnection()
	if connection == nil {
		t.Error("could not connect to database.")
	}

	// connection.CreateTable(&Resource{})
	defer connection.Close()
}

func TestDelete(t *testing.T) {
	connection := GetConnection()
	if connection == nil {
		t.Error("could not connect to database.")
	}
	connection.Where("from = ?", "https://huotu42.com/").Find(&Resource{})
	// connection.Unscoped().Where("from = https://huotu42.com/").Delete(&Resource{})

	// connection.CreateTable(&Resource{})
	defer connection.Close()
}
