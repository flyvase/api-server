package main

import (
	"log"
	"net/http"
	"os"

	"harvest/handlers"
	"harvest/repositories"
)

func main() {
	log.SetFlags(0)

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
