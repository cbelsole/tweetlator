package tweetlator

import (
	"encoding/json"
	"net/http"
)

func jsonResponse(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func jsonError(w http.ResponseWriter, r *http.Request, err error, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
}
