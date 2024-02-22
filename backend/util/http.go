package util

import (
	"encoding/json"
	"net/http"
)

func ReadJSON[T any](r *http.Request) (T, error) {
	decoder := json.NewDecoder(r.Body)
	var t T
	err := decoder.Decode(&t)
	return t, err
}

func WriteJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
