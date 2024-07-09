package router

import (
	"github.com/gorilla/mux"
	"github.com/syrus/mangoapi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movies", controller.CreateMovies).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controller.DeleteAmovie).Methods("DELETE")
	router.HandleFunc("/api/movies/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/dleteallmovie", controller.DeleteAllmovie).Methods("DELETE")
	return router
}
