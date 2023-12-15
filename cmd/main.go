package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")

		responseWriter.WriteHeader(http.StatusOK)

		err := json.NewEncoder(responseWriter).Encode(struct {
			Status string `json:"status"`
		}{
			Status: "OK",
		})
		if err != nil {
			http.Error(responseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	})

	log.Printf("Starting http server on :8000\n")
	log.Printf("Open http://localhost:8000/health to check status\n")

	http.ListenAndServe(":8000", nil)
}
