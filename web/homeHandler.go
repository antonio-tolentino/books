package web

import (
	"encoding/json"
	"net/http"
)

// Home Handler
func homeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("Books API")
}
