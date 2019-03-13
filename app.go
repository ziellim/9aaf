package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"log"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/v0/landmarks", a.getlandmarks).Methods("GET")
}

func (a *App) getlandmarks(w http.ResponseWriter, r *http.Request) {
	tm := Landmark{}
	landmarks, err := tm.getAll()
	if err != nil {
		respondWithJSON(w, http.StatusOK, landmarks)
	}
	respondWithError(w, http.StatusInternalServerError, err.Error())
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}