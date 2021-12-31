package main

import (
	"api-server/src/application/repository"
	"api-server/src/config"
	"api-server/src/infrastructure/gateway"
	"api-server/src/infrastructure/gateway/firebase"
	"api-server/src/infrastructure/gateway/sql"
	"api-server/src/infrastructure/http/handler"
	"api-server/src/infrastructure/http/middleware"
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
	firebaseAuth := gateway.InitializeAuth(firebaseApp)
	firebaseAuthImpl := gateway.AuthImpl{
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

	mux.Handle("/spaces/{space_id:[0-9]{1,10}}/",
		middleware.Defaults(
			&authRepository,
			handler.SpaceDetailsGet(
				&spaceRepository,
			),
		),
	).Methods("GET")

	mux.Handle("/spaces/{space_id:[0-9]{1,10}}/ogp/",
		middleware.Defaults(
			&authRepository,
			handler.SpaceOgpGet(
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
