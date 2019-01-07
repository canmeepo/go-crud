package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id    int    `json:id`
	Title string `json:title`
	Genre string `json:genre`
	Year  string `json:year`
}

var movies []Movie

func main() {
	router := mux.NewRouter()

	movies = append(movies, Movie{Id: 1, Title: "avatar", Genre: "action", Year: "2009"},
		Movie{Id: 2, Title: "the avengers", Genre: "action", Year: "2012"})

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
	params := mux.Vars(r)

	i, _ := strconv.Atoi(params["id"])

	for _, movie := range movies {
		if movie.Id == i {
			json.NewEncoder(w).Encode(&movie)
		}
	}
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie

	json.NewDecoder(r.Body).Decode(&movie)

	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie

	json.NewDecoder(r.Body).Decode(&movie)

	for i, item := range movies {
		if item.Id == movie.Id {
			movies[i] = movie
		}
	}

	json.NewEncoder(w).Encode(&movie)
}

func removeMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	for i, item := range movies {
		if item.Id == id {
			movies = append(movies[:i], movies[i+1:]...)
		}
	}

	json.NewEncoder(w).Encode(movies)

}
