package main

import (
	"encoding/json"
	"net/http"
)

type Coaster struct {
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	ID           string `json:"id"`
	InPark       string `json:"inPark"`
	Height       int    `json:"height"`
}
type coasterHandlers struct {
	store map[string]Coaster
}

func newCoasterHandlers() *coasterHandlers {
	return &coasterHandlers{
		store: map[string]Coaster{
			"Id1": Coaster{
				Name:         "Maneesh Kumar",
				Manufacturer: "B+m",
				ID:           "Id1",
				InPark:       "Caro Land",
				Height:       995,
			},
		},
	}

}

//Both parameter is the interface (w http.ResponseWriter, r *http.Request)
func (h *coasterHandlers) get(w http.ResponseWriter, r *http.Request) {
	coasters := make([]Coaster, len(h.store))

	i := 0
	for _, coaster := range h.store {
		coasters[i] = coaster
		i++
	}

	//CONVERT A LIST INTO JSON
	jsonBytes, err := json.Marshal(coasters)
	if err != nil {
		//To do
	}

	w.Write(jsonBytes)
}

func main() {
	// Register the function or define handler function
	coasterHandlers := newCoasterHandlers()
	http.HandleFunc("/coasters", coasterHandlers.get)
	// Creating a Simple http Server
	// Port , default Handler=nil
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
