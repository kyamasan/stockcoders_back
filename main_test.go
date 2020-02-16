package main

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(t *testing.T) {
	//Dotenvを使用
	godotenv.Load(".env.test")
	env := os.Getenv("DB_ENV")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("SERVER_PORT")

	//.envファイルが読み込めていない
	if env != "TEST_ENV" {
		t.Fatalf("Wrong ENV, got %q", env)
	}
	if name != "TEST_NAME" {
		t.Fatalf("Wrong NAME, got %q", name)
	}
	if user != "TEST_USER" {
		t.Fatalf("Wrong USER, got %q", user)
	}
	if password != "TEST_PASSWORD" {
		t.Fatalf("Wrong PASSWORD, got %q", password)
	}
	if port != "TEST_PORT" {
		t.Fatalf("Wrong PORT, got %q", port)
	}
}
