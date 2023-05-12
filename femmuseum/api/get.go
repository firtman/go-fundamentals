package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"frontendmuseum.com/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	id := queryString["id"]
	if id != nil {
		index, err := strconv.Atoi(id[0])
		if err == nil && index < len(data.GetAll()) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data.GetAll()[index])
		} else {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data.GetAll())
	}
}
