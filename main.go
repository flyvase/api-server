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
	"harvest/infrastructure/http/handler"
	"harvest/infrastructure/http/middleware"
	"harvest/infrastructure/repositoryimpl"
	"harvest/infrastructure/sql"
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

	sql := sql.NewSqlImpl()
	userRepo := &repositoryimpl.User{Sql: sql}
	authRepo := &repositoryimpl.Auth{Client: auth}
	spaceRepo := &repositoryimpl.Space{Sql: sql}

	mux := http.NewServeMux()
	mux.Handle("/user/", middleware.Demux(
		&middleware.Group{
			Post: handler.UserPost(authRepo, userRepo),
		},
	))
	mux.Handle("/space/", middleware.Demux(
		&middleware.Group{
			Get: handler.SpaceGet(authRepo, spaceRepo),
		},
	))

	c := cors.New(cors.Options{
		AllowedOrigins: config.AllowedOrigin(),
		AllowedMethods: []string{http.MethodPost, http.MethodGet},
		AllowedHeaders: []string{"Authorization", "Content-Type", "X-Cloud-Trace-Context"},
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(":"+port, c.Handler(mux)); err != nil {
		panic(err)
	}
}
