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
		// fmt.Println("World")
		// fmt.Println(r.Method) -> Show HTTP Methode Used
		// w.Write([]byte("Hai")) -> Send Response
		if r.Method == http.MethodGet {
			//Data Response
			data := responses.Response{
				Code: http.StatusOK,
				Body: "Pong",
			}
			//Make JSON Respones
			json.NewEncoder(w).Encode(data)
		}
	})
	//Start Server
	http.ListenAndServe("localhost:3000", mux)
}
