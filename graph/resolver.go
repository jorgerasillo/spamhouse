package graph

import (
	"github.com/jorgerasillo/spamhouse/repo"
	"github.com/jorgerasillo/spamhouse/repo/db"
	"github.com/sirupsen/logrus"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	Repository repo.Repository
	Logger     *logrus.Logger
	QChan      chan *db.IPAddress
}
