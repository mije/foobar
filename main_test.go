package main

import (
	"log"
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	dbURL := os.Getenv("DB_URL")
	log.Printf("DB_URL=%q", dbURL)
}
