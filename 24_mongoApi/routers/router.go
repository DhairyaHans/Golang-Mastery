package routers

import (
	"github.com/dhairyahans/mongoapi/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	r.HandleFunc("/api/movie", controllers.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movie/{id}", controllers.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/api/movie/{id}", controllers.DeleteOneMovie).Methods("DELETE")
	r.HandleFunc("/api/movies", controllers.DeleteAllMovies).Methods("DELETE")

	return r
}
