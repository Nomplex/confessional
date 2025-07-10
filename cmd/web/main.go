package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/nomplex/confessional/internal/models"
)

type application struct {
	logger      *slog.Logger
	confessions *models.ConfessionModel
}

func main() {
	port := flag.String("port", "4242", "port to use")
	fresh_database := flag.Bool("fresh-database", false, "Use a fresh database")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	err := godotenv.Load()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	dsn := os.Getenv("DSN")
	db, err := openDB(dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	if *fresh_database {
		// Setup database
		initDB(db)
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}

		logger.Info("Database Reset")
	}

	app := &application{
		logger:      logger,
		confessions: &models.ConfessionModel{DB: db},
	}
	srv := &http.Server{
		Addr:    ":" + *port,
		Handler: app.router(),
	}

	logger.Info("Starting server", "addr", srv.Addr)

	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func initDB(db *sql.DB) error {
	seedFile, err := os.ReadFile("./seed.sql")
	if err != nil {
		return err
	}

	stmts := strings.SplitSeq(string(seedFile), ";")

	for stmt := range stmts {

		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}

		_, err = db.Exec(stmt)
		if err != nil {
			return err
		}
	}

	return nil
}
