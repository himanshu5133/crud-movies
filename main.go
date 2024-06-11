package main

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}
type Director struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:= mux.Vars(r)

	for _,item:= range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:= mux.Vars(r)

	for index,item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:= mux.Vars(r)

	for index,item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Intn(100000000))
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
}

var movies []Movie

func main(){
	r:=mux.NewRouter()
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Server starting at port :8000\n")
	
	if err:=http.ListenAndServe(":8000",r);err!=nil{
		log.Fatal(err)
	}
}