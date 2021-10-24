package Handle

import "github.com/gorilla/mux"

func NewServer(handler FCHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/flashcard/{Type}", handler.CreateCards).Methods("POST")

	return r
}