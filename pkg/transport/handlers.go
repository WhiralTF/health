package transport

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/health", Health).Methods(http.MethodHead)
	return r
}

func Health(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "it's working!")
}
