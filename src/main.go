package main

import (
	"fmt"
	"harvest/src/adapter/gateway/sql"
	"harvest/src/adapter/http/handler"
	"harvest/src/adapter/http/middleware"
	"harvest/src/application/repository"
	"harvest/src/config"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	sqlDriver := sql.NewDriver()
	spaceRepository := repository.SpaceImpl{
		SqlDriver: sqlDriver,
	}

	mux := mux.NewRouter()
	mux.Handle(
		"/spaces/",
		middleware.Logger(
			handler.SpacesGet(&spaceRepository),
		),
	).Methods("GET")
	mux.Handle("/spaces/{space_id:[0-9]{1,10}}",
		handler.SpaceDetailsGet(
			&spaceRepository,
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
		fmt.Println("Starting web server")
	}
	if err := http.ListenAndServe(":"+port, c.Handler(mux)); err != nil {
		panic(err)
	}
}
