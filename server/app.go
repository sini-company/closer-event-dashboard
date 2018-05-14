package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Event closer webhook delivery payload
type Event struct {
	ID         string                 `json:"id"`
	Event      string                 `json:"event"`
	SourceID   string                 `json:"sourceId"`
	SourceType string                 `json:"sourceType"`
	Data       map[string]interface{} `json:"data"`
	Timestamp  int64                  `json:"timestamp"`
}

// Delivery closer webhook delivery payload
type Delivery struct {
	ID         string  `json:"id"`
	WebhookID  string  `json:"webhookId"`
	WebhookURL string  `json:"webhookUrl"`
	Messages   []Event `json:"messages"`
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// WebhookServer returns the value for key or ok=false if there is no mapping for key.
func WebhookServer(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" && req.URL.Path == "/webhook-endpoint" {
		var delivery Delivery
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&delivery)
		defer req.Body.Close()
		if err != nil {
			http.Error(w, `400 bad request`, http.StatusBadRequest)
		}
		log.Printf("received %d event(s).", len(delivery.Messages))
		InsertEvents(delivery.Messages...)
		io.WriteString(w, "OK")

	} else if req.Method == "GET" && req.URL.Path == "/health" {
		io.WriteString(w, "OK")

	} else {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
}

func main() {
	envPort, _ := strconv.Atoi(getEnv("PORT", "3000"))
	port := flag.Int("port", envPort, "port of http server")
	flag.Parse()

	db := GetDBInstance()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(90) // max_connection is set to 100 (default)
	defer db.Close()

	http.HandleFunc("/", WebhookServer)
	log.Printf("start webhook server on port %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
