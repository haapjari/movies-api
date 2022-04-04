package main

import (
	"github.com/gorilla/mux"
)

// "fmt"
// "log"
// "encoding/json"
// "math/rand"
// "net/http"
// "strconv"
// "github.com/gorilla/mux"

type Movie struct {
	IO string `json: "id"`
	ISBN string `json: "isbn"`
	Title string `json: "title"`
	Director *Director `json: "director"`
}

type Director struct {
	FirstName string `json: "firstName"`
	LastName string `json: "lastName"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc()	
}