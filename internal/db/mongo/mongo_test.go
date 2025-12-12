package mongo

import (
	"os"
	"testing"
)

func TestConnect(t *testing.T) {
	uri := os.Getenv("MONGODB_URI")
	client, err := Connect(uri)
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer Disconnect(client)

	if client == nil {
		t.Fatal("Got nil client despite no error")
	}
}

func TestConnectWithInvalidURI(t *testing.T) {
	client, err := Connect("invalid")
	if err == nil {
		t.Error("Expected error with invalid URI")
	}
	if client != nil {
		t.Error("Expected nil client on error")
	}
}