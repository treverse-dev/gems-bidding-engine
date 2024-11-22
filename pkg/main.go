package main

import (
	"fmt"
	"net/http"
	_ "net/http"
	_ "os"
)

func main() {
	http.HandleFunc("/bid", handleBid)
	fmt.Println("Bidding service started on port 8080")
	http.ListenAndServe(":8080", nil)
}

func handleBid(w http.ResponseWriter, r *http.Request) {
	// Implement bidding logic here
	fmt.Fprintf(w, "Bid received")
}
