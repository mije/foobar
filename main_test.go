package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"

	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

func init() {
	mustGetenv := func(key string) string {
		val := os.Getenv(key)
		if val == "" {
			log.Fatalf("test: env: unable to find %q", key)
		}
		return val
	}
	m, err := migrate.New(mustGetenv("DB_MIGRATIONS"), mustGetenv("DB_URL"))
	if err != nil {
		log.Fatalf("test: create migration: %v", err)
	}

	if err := m.Up(); err != nil {
		log.Fatalf("test: database up: %v", err)
	}
}

func TestEnv(t *testing.T) {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		t.Fatal("unable to find \"DB_URL\"")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		t.Fatal(err)
	}
	for rows.Next() {
		var (
			id        int
			firstName string
			lastName  string
		)
		err = rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("%v | %v | %v\n", id, firstName, lastName)
	}
	if err := rows.Err(); err != nil {
		t.Fatal(err)
	}
}
