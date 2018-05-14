package main

import (
	"fmt"
	"net/http"
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	_, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s/health", getEnv("PORT", "3000")))
	if err != nil {
		os.Exit(1)
	}
}
