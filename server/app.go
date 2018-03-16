package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
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
		for _, element := range delivery.Messages {
			InsertEvent(element)
		}
		io.WriteString(w, "OK")
	} else {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
}

func main() {
	port := flag.Int("port", 3000, "port of http server")
	flag.Parse()

	http.HandleFunc("/", WebhookServer)
	log.Printf("start webhook server on port %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
