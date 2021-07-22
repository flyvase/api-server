package main

import (
	"net/http"
	"os"

	"harvest/config"
	"harvest/controllers"
	"harvest/entities"
	"harvest/handlers"
	"harvest/repositories"
)

func main() {
	if config.Mode == "release" {
		cpr := repositories.CloudProfiler{}
		if err := controllers.StartProfiler(cpr, entities.ProfilerConfig{NoCPUProfiling: true}); err != nil {
			panic("Failed to start profiling")
		}
	}

	mux := http.NewServeMux()
	mux.Handle("/users/", handlers.UsersHandler())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		panic(err)
	}
}
