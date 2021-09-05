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

func initializeFirebase() *firebase.App {
	if config.Mode == "debug" && config.Environment == "dev" {
		opt := option.WithCredentialsFile("keys/harvest-firebase-admin-sa-key.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			panic(err)
		}

		return app
	} else {
		app, err := firebase.NewApp(context.Background(), nil)
		if err != nil {
			panic(err)
		}

		return app
	}
}

func main() {
	log.SetFlags(0)
	logger.Debug("Starting web server", "main")

	app := initializeFirebase()
	auth, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}

	sqlRepo := repository.NewSqlRepositoryImpl()
	userRepo := &repository.UserImpl{Sql: sqlRepo}
	authRepo := &repository.AuthImpl{Client: auth}

	mux := http.NewServeMux()
	mux.Handle("/user/", handler.UserHandler(userRepo, authRepo))

	c := cors.New(cors.Options{
		AllowedOrigins: config.AllowedOrigin(),
		AllowedMethods: []string{http.MethodPost},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(":"+port, c.Handler(mux)); err != nil {
		panic(err)
	}
}
