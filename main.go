package main

import (
	"log"
	"net/http"
	"os"

	"harvest/interfaces"
	"harvest/repositories"
	"harvest/web"
)

func main() {
	log.SetFlags(0)

	db, err1 := repositories.InitMySqlConnection()
	if err1 != nil {
		panic(err1)
	}

	provider := interfaces.RepositoriesProvider{
		User: repositories.User{DB: db},
	}

	mux := http.NewServeMux()
	mux.Handle("/users/", web.UsersHandler(&provider))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err2 := http.ListenAndServe(":"+port, mux); err2 != nil {
		panic(err2)
	}
}
