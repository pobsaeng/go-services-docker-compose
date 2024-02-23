package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Define routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/findall", findAllFromService2)

	// Start the server
	fmt.Println("Server is running on port 4000...")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func findAllFromService2(w http.ResponseWriter, r *http.Request) {
	// Make a GET request to service2
	resp, err := http.Get("http://service2:4001/findall")
	if err != nil {
		http.Error(w, "Error fetching data from service2", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy the response from service2 to the response writer
	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Error copying response body", http.StatusInternalServerError)
		return
	}
}
