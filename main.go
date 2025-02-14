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
	DB             *database.Queries
	authSecret     string
	polkaKey       string
}



func main() {
	mux := http.ServeMux{}
	server := http.Server{
		Handler: &mux,
		Addr:    ":8080",
	}

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	cfg := &apiConfig{}

	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}

	cfg.authSecret = os.Getenv("AUTH_SECRET")
	cfg.polkaKey = os.Getenv("POLKA_KEY")

	cfg.DB = database.New(db)

	mux.Handle("/app/", cfg.middlewareMetricsInc(http.StripPrefix("/app/", http.FileServer(http.Dir("app")))))

	mux.HandleFunc("GET /admin/metrics", cfg.handleMetrics)
	mux.HandleFunc("POST /admin/reset", cfg.handleReset)
	mux.HandleFunc("GET /api/healthz", handleHealth)

	mux.HandleFunc("POST /api/login", cfg.handleLogin)
	mux.HandleFunc("POST /api/refresh", cfg.handleRefresh)
	mux.HandleFunc("POST /api/revoke", cfg.handleRevoke)

	mux.HandleFunc("POST /api/chirps", cfg.handleCreateChirp)

	mux.HandleFunc("GET /api/chirps", cfg.handleChirps)
	mux.HandleFunc("GET /api/chirps/{chirpID}", cfg.handleGetChirp)
	mux.HandleFunc("DELETE /api/chirps/{chirpID}", cfg.handleDeleteChirp)


	mux.HandleFunc("POST /api/users", cfg.handleCreateUser)
	mux.HandleFunc("PUT /api/users", cfg.handlePutUser)

	mux.HandleFunc("POST /api/polka/webhooks", cfg.handlePolkaWebhook)


	println("Starting server on :8080")

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
