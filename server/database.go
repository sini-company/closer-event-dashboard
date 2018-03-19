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

// GetDBConnection get database connection
func GetDBConnection() *sql.DB {
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

// InsertEvent insert each event
func InsertEvent(event Event) {
	db := GetDBConnection()
	data, _ := json.Marshal(event.Data)
	dataString := string(data)

	if string(data) != "null" {
		dataString = fmt.Sprintf("'%s'", dataString)
	}

	query := fmt.Sprintf(
		oneLine(`
			INSERT INTO events (id, sourceId, sourceType, event, data, timestamp)
			VALUES ('%s', '%s', '%s', '%s', %s, '%s');
		`),
		event.ID,
		event.SourceID,
		event.SourceType,
		event.Event,
		dataString,
		time.Unix(event.Timestamp/1000, event.Timestamp%1000).UTC().Format(time.RFC3339))

	_, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
}
