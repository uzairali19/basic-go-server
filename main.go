package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// User - a simple struct to represent a User
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// users - a slice to store users (in lieu of a database)
var users = []User{
	{ID: 1, Name: "John Doe"},
	{ID: 2, Name: "Jane Doe"},
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/users", getUsers)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
