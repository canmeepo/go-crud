package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json" 
)

type Movie struct {
	Id	int	`json:id`
	Title string `json:title`
	Genre string  `json:genre`
	Year string `json:year`
}

var movies []Movie

func main() {
	router:= mux.NewRouter()

	movies = append(movies, Movie{Id: 1, Title: "avatar", Genre: "action", Year: "2009"})

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", addMovie).Methods("POST")
	router.HandleFunc("/movies", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", removeMovie).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	log.Println("get one movie")
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	log.Println("add movie")
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	log.Println("update movie")
}

func removeMovie(w http.ResponseWriter, r *http.Request) {
	log.Println("remove movie")
}