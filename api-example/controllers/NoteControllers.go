package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/aslampangestu/learn-golang/api-example/views"

	"github.com/aslampangestu/learn-golang/api-example/models"
)

// IndexNote -> Create & Select All & Select By Query Data
func IndexNote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Insert Data
		if r.Method == http.MethodPost {
			// Take some data
			data := views.PostRequestNote{}
			json.NewDecoder(r.Body).Decode(&data)
			// Save it
			if err := models.InsertNote(data.Title, data.Desc); err != nil {
				w.Write([]byte("Some Error"))
				return
			}
			res := views.ResponseWithData{
				Code:    http.StatusCreated,
				Status:  true,
				Message: "Succes Create Note",
				Data:    data,
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(res)
			//Select All Data
		} else if r.Method == http.MethodGet {
			title := r.URL.Query().Get("title")
			desc := r.URL.Query().Get("desctiption")
			// id := r.URL.Path[1:]
			if title == "" && desc == "" {
				data, err := models.SelectAllNote()
				SetReturnGet(w, data, err)
			} else {
				data, err := models.SelectValueNote(title, desc)
				SetReturnGet(w, data, err)
			}
		}
	}
}

// SetReturnGet -> Return Get
func SetReturnGet(w http.ResponseWriter, data []views.GetRequestNote, err error) {
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
