package main

import (
	"context"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go/v4"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"google.golang.org/api/option"

	"harvest/config"
	"harvest/core/logger"
	"harvest/infrastructure/http/handler"
	"harvest/infrastructure/http/middleware"
	"harvest/infrastructure/repository"
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

	driver := sql.NewDriverImpl()
	userRepo := &repository.User{Driver: driver}
	authRepo := &repository.Auth{Client: auth}
	spaceRepo := &repository.Space{Driver: driver}

	mux := mux.NewRouter()
	mux.Handle("/users/", middleware.Demux(&middleware.Group{Post: handler.UsersPost(authRepo, userRepo)}))
	mux.Handle("/spaces/", middleware.Demux(&middleware.Group{Get: handler.SpacesGet(authRepo, spaceRepo)}))
	mux.Handle("/spaces/{space_id:[0-9]{0,10}}/", middleware.Demux(&middleware.Group{Get: handler.SpaceDetailsGet(authRepo, spaceRepo)}))

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
