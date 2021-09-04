package main

import (
	"context"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go/v4"
	"github.com/rs/cors"
	"google.golang.org/api/option"

	"harvest/config"
	"harvest/core/logger"
	"harvest/infrastructure/repository"
	"harvest/infrastructure/web/handler"
)

func initializeFirebase() (*firebase.App, context.Context) {
	ctx := context.Background()

	if config.Mode == "debug" && config.Environment == "dev" {
		opt := option.WithCredentialsFile("keys/harvest-firebase-admin-sa-key.json")
		app, err := firebase.NewApp(ctx, nil, opt)
		if err != nil {
			panic(err)
		}

		return app, ctx
	} else {
		app, err := firebase.NewApp(ctx, nil)
		if err != nil {
			panic(err)
		}

		return app, ctx
	}
}

func main() {
	log.SetFlags(0)
	logger.Debug("Starting web server", "main")

	initializeFirebase()

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
