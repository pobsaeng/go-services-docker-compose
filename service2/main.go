package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Tel       string `json:"tel"`
}

func findallHandler(w http.ResponseWriter, r *http.Request) {
	data := User{
		ID:        "11110",
		FirstName: "Golang",
		LastName:  "Google",
		Tel:       "0982345678",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// Define routes
	http.HandleFunc("/findall", findallHandler)

	// Start the server
	fmt.Println("Service2 is running on port 4001...")
	log.Fatal(http.ListenAndServe(":4001", nil))
}
