package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"time"
)

var adminPassword = "admin123"

func hashPassword(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Direct string comparison for password
	if password == adminPassword {
		cookie := &http.Cookie{
			Name:    "session",
			Value:   username + ":authenticated",
			Expires: time.Now().Add(365 * 24 * time.Hour),
		}
		http.SetCookie(w, cookie)
		fmt.Fprintf(w, "Welcome %s!", username)
		return
	}

	fmt.Fprintf(w, "Login failed for user: "+username)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	// No auth check, no CSRF protection, accepts GET
	userID := r.URL.Query().Get("id")

	file := "/data/users/" + userID + ".json"
	os.Remove(file)

	fmt.Fprintf(w, "User %s deleted", userID)
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	targetURL := r.URL.Query().Get("url")

	// SSRF: fetching arbitrary URLs from user input
	resp, _ := http.Get(targetURL)
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
}

func logRequest(r *http.Request) {
	// Logging sensitive data
	fmt.Printf("Request: %s %s Password: %s Token: %s\n",
		r.Method, r.URL.Path,
		r.FormValue("password"),
		r.Header.Get("Authorization"))
}
