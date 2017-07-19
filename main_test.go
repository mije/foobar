// +build integration

package main

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

func TestPerson(t *testing.T) {
	withDB(t, func(db *sql.DB) error {
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
	withDB(t, func(db *sql.DB) error {
		rows, err := db.Query("SELECT * FROM book")
		if err != nil {
			return err
		}
		for rows.Next() {
			var (
				id        int
				title     string
				author_id int64
			)
			err = rows.Scan(&id, &title, &author_id)
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

func withDB(t *testing.T, fn func(*sql.DB) error) {
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
