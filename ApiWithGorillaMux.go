package main


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// type Article struct {
// 	Title   string `json:"Title"`
// 	Desc    string `json:"desc"`
// 	Content string `json:"content"`
// }

// // Existing code from above
// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the HomePage!")
// 	fmt.Println("Endpoint Hit: homePage")
// }

// var Articles []Article

// func returnAllArticles(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: returnAllArticles")
// 	json.NewEncoder(w).Encode(Articles)
// }
// func handleRequests() {
// 	// creates a new instance of a mux router
// 	myRouter := mux.NewRouter().StrictSlash(true)
// 	// replace http.HandleFunc with myRouter.HandleFunc
// 	myRouter.HandleFunc("/", homePage)
// 	myRouter.HandleFunc("/all", returnAllArticles)

// 	// finally, instead of passing in nil, we want
// 	// to pass in our newly created router as the second
// 	// argument
// 	log.Fatal(http.ListenAndServe(":10000", myRouter))
// }

// func main() {
// 	fmt.Println("Rest API v2.0 - Mux Routers")
// 	Articles = []Article{
// 		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
// 		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
// 	}
// 	handleRequests()
// }

// //(2)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Existing code from above
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

var Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}
func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	// FOR RETURNING SINGLE ARTICLE
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

// func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	key := vars["id"]

// 	// fmt.Fprintf(w, "Key: "+key)
// 	// fmt.Fprintf(w, "Key: ", vars)

// }
func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// Loop over all of our Articles
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}
func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
