package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type application struct{}

func main() {
	err := godotenv.Load()
	if err != nil {
		os.Exit(1)
	}

	port := os.Getenv("PORT")

	app := &application{}
	srv := &http.Server{
		Addr:    port,
		Handler: app.router(),
	}

	fmt.Printf("Starting server on %s\n", port)

	err = srv.ListenAndServe()
	os.Exit(1)
}
