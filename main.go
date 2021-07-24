package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/profiler"
	"github.com/pkg/errors"

	"harvest/config"
	"harvest/handlers"
	"harvest/repositories"
)

func main() {
	log.SetFlags(0)

	if config.Mode == "release" {
		if err := profiler.Start(profiler.Config{NoCPUProfiling: true}); err != nil {
			msg := fmt.Sprintf("Failed to start profiler\n%+v", errors.WithStack(err))
			panic(msg)
		}
	}

	db, err1 := repositories.InitMySqlConnection()
	if err1 != nil {
		panic(err1)
	}

	ur := repositories.User{db}

	mux := http.NewServeMux()
	mux.Handle("/users/", handlers.UsersHandler(ur))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err2 := http.ListenAndServe(":"+port, mux); err2 != nil {
		panic(err2)
	}
}
