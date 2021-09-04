package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"

	"harvest/config"
	"harvest/core/logger"
	"harvest/infrastructure/repository"
	"harvest/infrastructure/web/handler"
)

func main() {
	log.SetFlags(0)
	logger.Debug("Starting web server", "main")

	sqlRepo := repository.NewSqlRepositoryImpl()
	userRepo := repository.UserImpl{Sql: sqlRepo}

	mux := http.NewServeMux()
	mux.Handle("/user/", handler.UserHandler(userRepo))

	cors.New(cors.Options{
		AllowedOrigins: config.AllowedOrigin(),
		AllowedMethods: []string{http.MethodPost},
	}).Handler(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		panic(err)
	}
}
