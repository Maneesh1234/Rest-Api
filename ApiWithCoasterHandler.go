// package main
//go run SecondApi.go
// cd "Restful Api"

// import (
// 	"encoding/json"
// 	"net/http"
// )

// type Coaster struct {
// 	Name         string `json:"name"`
// 	Manufacturer string `json:"manufacturer"`
// 	ID           string `json:"id"`
// 	InPark       string `json:"inPark"`
// 	Height       int    `json:"height"`
// }
// type coasterHandlers struct {
// 	store map[string]Coaster
// }

// func newCoasterHandlers() *coasterHandlers {
// 	return &coasterHandlers{
// 		store: map[string]Coaster{
// 			"Id1": Coaster{
// 				Name:         "Maneesh Kumar",
// 				Manufacturer: "B+m",
// 				ID:           "Id1",
// 				InPark:       "Caro Land",
// 				Height:       995,
// 			},
// 		},
// 	}

// }

// //Both parameter is the interface (w http.ResponseWriter, r *http.Request)
// func (h *coasterHandlers) get(w http.ResponseWriter, r *http.Request) {
// 	coasters := make([]Coaster, len(h.store))

// 	i := 0
// 	for _, coaster := range h.store {
// 		coasters[i] = coaster
// 		i++
// 	}

// 	jsonBytes, err := json.Marshal(coasters)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(err.Error()))
// 	}
// 	w.Header().Add("content-type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(jsonBytes)
// }

// func main() {
// 	// Register the function or define handler function
// 	coasterHandlers := newCoasterHandlers()
// 	http.HandleFunc("/coasters", coasterHandlers.get)
// 	// Creating a Simple http Server
// 	// Port , default Handler=nil
// 	err := http.ListenAndServe(":8081", nil)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// (2)
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Coaster struct {
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	ID           string `json:"id"`
	InPark       string `json:"inPark"`
	Height       int    `json:"height"`
}
type coasterHandlers struct {
	sync.Mutex
	store map[string]Coaster
}

// "Id1": Coaster{
// 	Name:         "Maneesh Kumar",
// 	Manufacturer: "B+m",
// 	ID:           "Id1",
// 	InPark:       "Caro Land",
// 	Height:       995,
// },
func newCoasterHandlers() *coasterHandlers {
	return &coasterHandlers{
		store: map[string]Coaster{},
	}

}

//ResponseWriter is the interface and  r *http.Request is a object reference
func (h *coasterHandlers) coasters(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

//Both parameter is the interface (w http.ResponseWriter, r *http.Request)
func (h *coasterHandlers) get(w http.ResponseWriter, r *http.Request) {
	coasters := make([]Coaster, len(h.store))

	i := 0
	h.Lock()
	// CONVERT IN FORM OF LIST OF COASTER
	for _, coaster := range h.store {
		coasters[i] = coaster
		i++
	}
	h.Unlock()
	// COVERT LIST INTO JSON
	jsonBytes, err := json.Marshal(coasters)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
func (h *coasterHandlers) post(w http.ResponseWriter, r *http.Request) {

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// IT  IS FOR RECEIVING DATA ONLY JSON FORMAT
	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("need content-type 'application/json', but got '%s'", ct)))
		return
	}

	var coaster Coaster
	//USER SEND DATA IN JSON FORMAT
	err = json.Unmarshal(bodyBytes, &coaster)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	coaster.ID = fmt.Sprintf("%d", time.Now().UnixNano())
	h.Lock()
	h.store[coaster.ID] = coaster
	defer h.Unlock()
}
func (h *coasterHandlers) getCoaster(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.URL)
	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// fmt.Println(parts[2])

	// if parts[2] == "random" {
	// 	h.getRandomCoaster(w, r)
	// 	return
	// }

	h.Lock()
	coaster, ok := h.store[parts[2]]
	h.Unlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonBytes, err := json.Marshal(coaster)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

// type adminPortal struct {
// 	password string
// }

// func newAdminPortal() *adminPortal {
// 	password := os.Getenv("ADMIN_PASSWORD")
// 	if password == "" {
// 		panic("required env var ADMIN_PASSWORD not set")
// 	}

// 	return &adminPortal{password: password}
// }

// func (a adminPortal) handler(w http.ResponseWriter, r *http.Request) {
// 	user, pass, ok := r.BasicAuth()
// 	if !ok || user != "admin" || pass != a.password {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		w.Write([]byte("401 - unauthorized"))
// 		return
// 	}

// 	w.Write([]byte("<html><h1>Super secret admin portal</h1></html>"))
// }

func main() {
	// admin := newAdminPortal()
	// Register the function or define handler function
	coasterHandlers := newCoasterHandlers()
	http.HandleFunc("/coasters", coasterHandlers.coasters)
	http.HandleFunc("/coasters/", coasterHandlers.getCoaster)
	// http.HandleFunc("/admin", admin.handler)
	// Creating a Simple http Server
	// Port , default Handler=nil
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
