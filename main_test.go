// +build integration

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

func withDB(t *testing.T, fn func(sql.DB) error) {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		t.Fatal("unable to find \"DB_URL\"")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := fn(db); err != nil {
		t.Fatal(err)
	}

	if err := db.Close(); err != nil {
		t.Log(err)
	}
}

func TestPerson(t *testing.T) {
	withDB(t, func(db sql.DB) error {
		rows, err := db.Query("SELECT * FROM person")
		if err != nil {
			return err
		}
		for rows.Next() {
			var (
				id        int
				firstName string
				lastName  string
			)
			err = rows.Scan(&id, &firstName, &lastName)
			if err != nil {
				return err
			}
			fmt.Printf("%v | %v | %v\n", id, firstName, lastName)
		}
		if err := rows.Err(); err != nil {
			return err
		}
		return nil
	})
}

func TestBook(t *testing.T) {
	withDB(t, func(db sql.DB) error {
		rows, err := db.Query("SELECT * FROM book")
		if err != nil {
			return err
		}
		for rows.Next() {
			var (
				id    int
				title string
			)
			err = rows.Scan(&id, &title)
			if err != nil {
				return err
			}
			fmt.Printf("%v | %v \n", id, title)
		}
		if err := rows.Err(); err != nil {
			return err
		}
		return nil
	})
}
