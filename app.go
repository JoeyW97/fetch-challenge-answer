package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	// Write the response body
	fmt.Fprintf(w, "Welcome to Fetch Receipt processor!")
}

func processReceipt(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var newReceipt ReceiptDto

		err := json.NewDecoder(r.Body).Decode(&newReceipt)
		if err != nil {
			http.Error(w, "The receipt is invalid", http.StatusBadRequest)
			return
		}

		receipt, err := newReceipt.toReceipt()
		if err != nil {
			http.Error(w, "The receipt is invalid", http.StatusBadRequest)
			return
		}
		id := addReceipt(receipt)
		response, err := json.Marshal(map[string]string{
			"id": id,
		})
		if err != nil {
			http.Error(w, "Error when creating response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Description", "Returns the ID assigned to the receipt")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(response)
		if err != nil {
			http.Error(w, "Error when writing response", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "HTTP method not supported", http.StatusMethodNotAllowed)
		return
	}

}

func getReceiptPoint(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		path := r.URL.Path
		components := strings.Split(path, "/")
		id := components[len(components)-2]
		receipt, found := getReceiptById(id)
		if !found {
			http.Error(w, "Receipt not found", http.StatusBadRequest)
			return
		}
		points := calculateAwardPoint(receipt)
		response, err := json.Marshal(map[string]int{
			"points": points,
		})
		if err != nil {
			http.Error(w, "Error when creating response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Description", "The number of points awarded")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(response)
		if err != nil {
			http.Error(w, "Error when writing response", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "HTTP method not supported", http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/receipts/{id}/points", getReceiptPoint).Methods("GET")
	router.HandleFunc("/receipts/process", processReceipt).Methods("POST")
	fmt.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
