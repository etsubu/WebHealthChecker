package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type HealthCheckResponse struct {
	Name string
	Version string
	Time int64
}

const version = "v1.0"
const applicationName = "WebHealthChecker"

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Content-Type", "application/json")
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func healthcheckResponder(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
	response, _ := json.Marshal(HealthCheckResponse{applicationName, version, time.Now().Unix()})
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/results", get).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/", post).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/", put).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/", delete).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/healthcheck", healthcheckResponder)
	r.HandleFunc("/api/v1/", notFound)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	log.Print("Starting http server")
	log.Fatal(http.ListenAndServe(":8080", r))
}