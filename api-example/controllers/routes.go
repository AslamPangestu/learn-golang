package controllers

import (
	"net/http"
)

// Register -> Gateway API
func Register() *http.ServeMux {
	mux := http.NewServeMux()
	//List Route
	mux.HandleFunc("/ping", Ping())
	mux.HandleFunc("/api/v1/notes", IndexNote())
	return mux
}
