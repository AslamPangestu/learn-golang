package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/aslampangestu/learn-golang/api-example/views"
)

// Ping -> Function Ping
func Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			//Data
			data := views.Response{
				Code:    http.StatusOK,
				Status:  true,
				Message: "Pong",
			}
			//Set Header
			w.WriteHeader(http.StatusOK)
			//Convert To JSON
			json.NewEncoder(w).Encode(data)
		}
	}
}
