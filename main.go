package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rs/cors"

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

	handlers.RegisterIndex()

	mux := http.DefaultServeMux
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	}).Handler(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Listening on port %s\n", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		panic(err)
	}
}
