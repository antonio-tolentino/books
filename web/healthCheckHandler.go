package web

import (
	"encoding/json"
	"log"
	"net/http"
)

type Health struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	log.Println("")

	health := Health{
		Name:   "Books API",
		Status: "UP",
	}

	json.NewEncoder(w).Encode(health)
}
