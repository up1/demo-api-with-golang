package handlers

import(
  "github.com/gorilla/mux"
)

func Router() *mux.Router {
  r := mux.NewRouter()
  r.HandleFunc("/home", home).Methods("GET")
  r.HandleFunc("/healthz", healthz)
  return r
}
