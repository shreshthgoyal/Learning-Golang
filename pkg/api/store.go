package api

import (
	"encoding/json"
	"net/http"
)

func (a *API) handleGet(w http.ResponseWriter, r *http.Request) {

	var inp struct {
		Key   string `json:""`
		Value string `json:""`
	}

	err := json.NewDecoder(r.Body).Decode(&inp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(a.str.Get(inp.Key)))
}

func (a *API) handleSet(w http.ResponseWriter, r *http.Request) {
	var inp struct {
		Key   string `json:""`
		Value string `json:""`
	}

	err := json.NewDecoder(r.Body).Decode(&inp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	a.str.Set(inp.Key, inp.Value)

	w.Write([]byte("key-val stored!\n"))
}
