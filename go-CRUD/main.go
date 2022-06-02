package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string     `json:"id"`
	Isbn     string     `json:"isbn"`
	title    string     `json:"title"`
	Director *Direactor `json:"director"`
}

type Direactor struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "4021132", title: "Movie Owner", Director: &Direactor{Firstname: "JON", Lastname: "doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "4021133", title: "Movie Owner", Director: &Direactor{Firstname: "JON", Lastname: "doe"}})
	r.HandleFunc("/movies", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(movies)
	}).Methods("GET")
	r.HandleFunc("/movies/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		for _, item := range movies {
			if item.ID == params["id"] {
				json.NewEncoder(w).Encode(item)
				return
			}
		}
	}).Methods("GET")
	r.HandleFunc("/movies", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		var movie Movie
		_ = json.NewDecoder(r.Body).Decode(&movie)
		movie.ID = strconv.Itoa(rand.Intn(1000000000))
		json.NewEncoder(w).Encode(movie)
	}).Methods("POST")
	r.HandleFunc("/movies/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		for index, item := range movies {
			if item.Isbn == params["id"] {
				movies = append(movies[:index], movies[index+1:]...)
				var movie Movie
				_ = json.NewDecoder(r.Body).Decode(&movie)
				movie.ID = params["id"]
				movies = append(movies, movie)
				json.NewEncoder(w).Encode(movie)
				return
			}
		}
	}).Methods("PUT")
	r.HandleFunc("/movies/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		for index, item := range movies {
			if item.ID == params["id"] {
				movies = append(movies[:index], movies[index+1:]...)
				break
			}
		}
	}).Methods("DELETE")

	fmt.Printf("starting server at port 3000")
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", r))
}
