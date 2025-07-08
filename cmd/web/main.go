package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type application struct {
	logger *slog.Logger
}

func main() {
	port := flag.String("port", "4242", "port to use")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	err := godotenv.Load()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		logger: logger,
	}
	srv := &http.Server{
		Addr:    ":" + *port,
		Handler: app.router(),
	}

	fmt.Printf("Starting server on %s\n", *port)

	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
