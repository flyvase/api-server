package main

import (
	"harvest/src/adapter/gateway/firebase"
	"harvest/src/adapter/gateway/sql"
	"harvest/src/adapter/http/handler"
	"harvest/src/adapter/http/middleware"
	"harvest/src/application/repository"
	"harvest/src/config"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	log.SetFlags(0)

	sqlDriver := sql.NewDriver()
	firebaseApp := firebase.InitializeApp()
	firebaseAuth := firebase.InitializeAuth(firebaseApp)
	firebaseAuthImpl := firebase.AuthImpl{
		Client: firebaseAuth,
	}
	authRepository := repository.AuthImpl{
		Client: &firebaseAuthImpl,
	}
	spaceRepository := repository.SpaceImpl{
		SqlDriver: sqlDriver,
	}

	mux := mux.NewRouter()

	mux.Handle(
		"/spaces/",
		middleware.Defaults(
			&authRepository,
			handler.SpacesGet(
				&spaceRepository,
			),
		),
	).Methods("GET")

	mux.Handle("/spaces/{space_id:[0-9]{1,10}}",
		middleware.Defaults(
			&authRepository,
			handler.SpaceDetailsGet(
				&spaceRepository,
			),
		),
	).Methods("GET")

	c := cors.New(
		cors.Options{
			AllowedOrigins: config.AllowedOrigin(),
			AllowedMethods: []string{http.MethodPost, http.MethodGet},
			AllowedHeaders: []string{"Authorization", "Content-Type", "X-Cloud-Trace-Context"},
		},
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if config.Mode == "debug" {
		log.Println("Starting web server")
	}
	if err := http.ListenAndServe(":"+port, c.Handler(mux)); err != nil {
		panic(err)
	}
}
