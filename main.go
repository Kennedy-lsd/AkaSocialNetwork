package main

import (
	"fmt"
	"log"
	"test/cmd/api"
	"test/env"
	"test/internal/data"
	"test/internal/db"
)

func main() {
	env := env.InitEnv()

	cfg := api.Config{
		Addr: fmt.Sprintf(":%s", env.Port),
		Db: api.DbConfig{
			Addr:         env.Dd_Addr,
			MaxOpenConns: env.DB_MAX_OPEN_CONNS,
			MaxIdleTime:  env.DB_MAX_IDLE_TIME,
		},
	}

	db, err := db.New(cfg.Db.Addr, 30, cfg.Db.MaxIdleTime)
	if err != nil {
		log.Fatalf("error with db connection %v", err)
	}

	defer db.Close()
	log.Println("database connection pool established")

	data := data.NewPostgresData(db)

	app := &api.App{
		Config: cfg,
		Data:   data,
	}

	mux := app.Mount()

	if err := app.Run(mux); err != nil {
		log.Fatalf("Application encountered an error: %v", err)
	}
}
