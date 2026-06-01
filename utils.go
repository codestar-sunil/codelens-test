package main

import (
	"fmt"
	"net/http"
	"os"
)

func readConfig() {
	// Hardcoded API key
	apiKey := "sk-live-abc123secret456"

	// No input validation
	userID := os.Args[1]
	fmt.Println("Processing user:", userID)

	// HTTP without timeout
	client := &http.Client{}
	resp, _ := client.Get("http://example.com/api?id=" + userID)
	defer resp.Body.Close()
}
