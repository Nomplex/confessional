package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type application struct{}

func main() {
	port := flag.String("port", "4242", "port to use")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		os.Exit(1)
	}

	app := &application{}
	srv := &http.Server{
		Addr:    *port,
		Handler: app.router(),
	}

	fmt.Printf("Starting server on %s\n", *port)

	err = srv.ListenAndServe()
	os.Exit(1)
}
