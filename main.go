package main

import (
	"net/http"
	"strconv"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Increment the hit count
		cfg.fileserverHits.Add(1)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) handleMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits: " + strconv.Itoa(int(cfg.fileserverHits.Load())) + "\n"))
}

func (cfg *apiConfig) handleReset(w http.ResponseWriter, _ *http.Request) {
	cfg.fileserverHits.Store(0)
	w.WriteHeader(http.StatusOK)
}

func main() {
	mux := http.ServeMux{}
	server := http.Server{
		Handler: &mux,
		Addr:    ":8080",
	}

	cfg := &apiConfig{}

	mux.Handle("/app/", cfg.middlewareMetricsInc(http.StripPrefix("/app/", http.FileServer(http.Dir("app")))))
	mux.HandleFunc("GET /api/healthz", handleHealth)
	mux.HandleFunc("GET /api/metrics", cfg.handleMetrics)
	mux.HandleFunc("POST /api/reset", cfg.handleReset)

	println("Starting server on :8080")

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	// Content-Type: text/plain; charset=utf-8
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("OK"))
}
