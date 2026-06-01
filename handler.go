package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os/exec"
)

func handleUser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	// SQL query with user input directly
	query := fmt.Sprintf("SELECT * FROM users WHERE name = '%s'", username)
	db, _ := sql.Open("mysql", "root:password@/mydb")
	db.Query(query)

	// Command injection
	cmd := exec.Command("bash", "-c", "echo "+username)
	output, _ := cmd.Output()

	w.Write(output)
}

func main() {
	http.HandleFunc("/user", handleUser)
	http.ListenAndServe(":8080", nil)
}
