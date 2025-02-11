package main

import (
	"database/sql"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/eefret/chirpy/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	DB *database.Queries
}



func main() {
	mux := http.ServeMux{}
	server := http.Server{
		Handler: &mux,
		Addr:    ":8080",
	}

	godotenv.Load()

	cfg := &apiConfig{}

	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}

	cfg.DB = database.New(db)

	mux.Handle("/app/", cfg.middlewareMetricsInc(http.StripPrefix("/app/", http.FileServer(http.Dir("app")))))

	mux.HandleFunc("GET /admin/metrics", cfg.handleMetrics)
	mux.HandleFunc("POST /admin/reset", cfg.handleReset)

	mux.HandleFunc("GET /api/healthz", handleHealth)
	mux.HandleFunc("POST /api/chirps", cfg.handleCreateChirp)
	mux.HandleFunc("POST /api/users", cfg.handleCreateUser)
	mux.HandleFunc("GET /api/chirps", cfg.handleChirps)
	mux.HandleFunc("GET /api/chirps/{chirpID}", cfg.handleGetChirp)
	mux.HandleFunc("POST /api/login", cfg.handleLogin)

	println("Starting server on :8080")

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
