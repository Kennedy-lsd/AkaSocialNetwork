package api

import (
	"log"
	"net/http"
	"test/internal/data"
	"time"
)

type Config struct {
	Addr string
	Db   DbConfig
}

type DbConfig struct {
	Addr         string
	MaxOpenConns string
	MaxIdleTime  string
}

type App struct {
	Config Config
	Data   data.Data
}

func (a *App) Mount() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/status", a.statusCheckHandler)
	mux.HandleFunc("POST /v1/users", a.createUserHandler)
	mux.HandleFunc("GET /v1/users", a.getUsersHandler)
	mux.HandleFunc("GET /v1/users/{id}", a.getUserHandler)
	mux.HandleFunc("DELETE /v1/users/{id}", a.deleteUserHandler)

	loggerMux := loggerMiddleware(mux)

	return loggerMux
}

func (a *App) Run(mux http.Handler) error {

	srv := &http.Server{
		Addr:         a.Config.Addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server is starting on %s", a.Config.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
		return err
	}

	return nil
}
