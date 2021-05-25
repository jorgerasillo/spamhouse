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
	"github.com/jorgerasillo/spamhouse/pkg/queue"
	"github.com/jorgerasillo/spamhouse/repo"
	"github.com/jorgerasillo/spamhouse/repo/db"
	"github.com/sirupsen/logrus"
)

func main() {
	// retrieve configuration
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Unable to retrieve configuration: %v\n", err)
	}

	// create connection to db and run auto migration
	database, err := db.GetDB(cfg.DataSourceName, cfg.DefaultUser, cfg.DefaultPassword)
	if err != nil {
		log.Fatalf("Unable to get db connection: %v\n", err)
	}

	// create repository
	repo, err := repo.New(database)
	if err != nil {
		log.Fatalf("Unable to create repository: %v\n", err)
	}

	log := logrus.New()
	setLogLevel(log, cfg.LogLevel)

	// startup queue
	qChan := make(chan *db.IPAddress)

	q := queue.New(repo, qChan, log)
	go q.Process()

	// create router and attach authz middleware

	authz := auth.New(repo)
	router := chi.NewRouter()
	router.Use(authz.Middleware())

	// startup graphql server
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Repository: repo, Logger: log, QChan: qChan}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", srv)
	router.Get("/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	log.Printf("connect to http://localhost:%s/ to access the GraphQL UI", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, router))
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
