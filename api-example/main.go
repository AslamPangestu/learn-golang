package main

import (
	"encoding/json"
	"net/http"

	"github.com/aslampangestu/learn-golang/api-example/responses"
)

func main() {
	mux := http.NewServeMux()
	//Route
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := responses.Response{
				Code: http.StatusOK,
				Body: "Pong",
			}
			json.NewEncoder(w).Encode(data)
		}
	})
	//Start Server
	http.ListenAndServe("localhost:3000", mux)
}
