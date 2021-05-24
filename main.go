package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/jorgerasillo/spamhouse/graph"
	"github.com/jorgerasillo/spamhouse/graph/generated"
	"github.com/jorgerasillo/spamhouse/pkg/config"
	"github.com/jorgerasillo/spamhouse/pkg/middleware/auth"
	"github.com/jorgerasillo/spamhouse/repo"
	"github.com/jorgerasillo/spamhouse/repo/db"
	"github.com/sirupsen/logrus"
)

const port = "8080"

func main() {
	// retrieve configuration
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Unable to retrieve configuration: %v\n", err)
	}

	// create connection to db and run auto migration
	db, err := db.GetDB(cfg.DataSourceName)
	if err != nil {
		log.Fatalf("Unable to get db connection: %v\n", err)
	}

	// create repository
	repo, err := repo.New(db)
	if err != nil {
		log.Fatalf("Unable to create repository: %v\n", err)
	}

	log := logrus.New()
	setLogLevel(log, cfg.LogLevel)

	// startup queue
	qChan := make(chan string)

	// create router and attach authz middleware
	router := chi.NewRouter()
	router.Use(auth.Middleware())

	// startup graphql server
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Repository: repo, Logger: log, QChan: qChan}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", srv)
	router.Get("/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func setLogLevel(log *logrus.Logger, level string) {
	switch level {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "severe":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}
}
