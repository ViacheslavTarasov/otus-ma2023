package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/health", HealthCheckHandler).Methods(http.MethodGet)
	router.Use(loggingMiddleware)
	http.Handle("/", router)

	srv := &http.Server{
		Handler:           router,
		Addr:              "0.0.0.0:8000",
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      15 * time.Second,
		ReadTimeout:       15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": http.StatusText(http.StatusOK)})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Println("Unable to encode JSON")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
