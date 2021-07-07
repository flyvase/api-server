package main

import (
	"log"
	"net/http"
	"os"

	"harvest/controllers"
	"harvest/entities"
	"harvest/handlers"
	"harvest/repositories"
)

func main() {
	if os.Getenv("MODE") == "release" {
		cpr := repositories.CloudProfiler{}
		if err := controllers.StartProfiler(cpr, entities.ProfilerConfig{NoCPUProfiling: true}); err != nil {
			panic("Failed to start profiling")
		}
	}

	handlers.RegisterIndex()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting port to %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
