package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"
	// postgres driver
	_ "github.com/lib/pq"
)

var once sync.Once
var instance *sql.DB

// oneLine trims all leading / traialing / doubled whitespace
func oneLine(str string) string {
	return strings.TrimSpace(regexp.MustCompile("(\\s{2,})").ReplaceAllString(str, " "))
}

// GetDBInstance get database
func GetDBInstance() *sql.DB {
	once.Do(func() {
		connStr := "postgres://postgres:postgres@db:5432/closer_event?sslmode=disable"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}
		instance = db
	})
	return instance
}

// InsertEvents insert each event
func InsertEvents(events ...Event) {
	db := GetDBInstance()
	stmt, err := db.Prepare(oneLine(`
		INSERT INTO events (id, source_id, source_type, event, data, timestamp)
		VALUES ($1, $2, $3, $4, $5, $6);
	`))
	if err != nil {
		log.Fatal(err)
	}

	for _, event := range events {
		data, _ := json.Marshal(event.Data)
		_, err = stmt.Exec(
			event.ID,
			event.SourceID,
			event.SourceType,
			event.Event,
			string(data),
			time.Unix(event.Timestamp/1000, event.Timestamp%1000).UTC().Format(time.RFC3339))

		if err != nil && !strings.Contains(err.Error(), "events_pkey") {
			fmt.Println(err)
		}
	}
}
