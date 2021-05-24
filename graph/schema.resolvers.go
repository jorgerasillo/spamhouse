package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"sync"

	"github.com/jorgerasillo/spamhouse/graph/generated"
	"github.com/jorgerasillo/spamhouse/graph/model"
	"github.com/jorgerasillo/spamhouse/pkg/spamhous"
	"github.com/jorgerasillo/spamhouse/repo/db"
	"github.com/sirupsen/logrus"
)

func (r *mutationResolver) Enqueue(ctx context.Context, input []string) (*model.IPAddressResult, error) {
	var wg sync.WaitGroup
	qChan := make(chan *db.IPAddress)

	for _, i := range input {
		ip, err := db.NewIP(i)
		r.Logger.WithFields(logrus.Fields{"ip": i, "uuid": ip.UUID}).Debug("processing input")

		if err != nil {
			r.Logger.WithFields(logrus.Fields{
				"err": err,
				"ip":  ip,
			}).Error("unable to create new ip address")
		}
		wg.Add(1)
		go spamhous.Query(&ip, qChan, &wg)
	}

	go func() {
		wg.Wait()
		close(qChan)
	}()

	addresses := make([]*model.IPAddress, 0)
	for val := range qChan {
		logEntry := r.Logger.WithFields(logrus.Fields{"ip": val.IP, "uuid": val.UUID})
		logEntry.Debug("ip received")

		// query if it exists
		_, err := r.Repository.GetIP(val.String())
		if err != nil {
			// check if not found
			if errors.Is(err, db.ErrRecordNotFound) {
				logEntry.Debug("record was not found, creating new entry")
				ipAddress, err := r.Repository.AddIP(*val)
				if err != nil {
					r.Logger.WithFields(logrus.Fields{
						"err": err,
						"ip":  val,
					}).Error("unable to add new ip address")
					continue
				}
				r.Logger.WithFields(logrus.Fields{
					"ip":   ipAddress.IP,
					"uuid": ipAddress.UUID}).Debug("new record added")

				addresses = append(addresses, modelToResponse(ipAddress))
			}
			logEntry.Debug("db error attempting to look for record")
			continue

		}

		logEntry.Debug("record existed, updating entry")
		// save if it does
		ipAddress, err := r.Repository.UpdateIP(*val)
		if err != nil {
			r.Logger.WithFields(logrus.Fields{
				"err": err,
				"ip":  val,
			}).Error("unable to update ip address")
			continue
		}
		r.Logger.WithFields(logrus.Fields{
			"ip":   ipAddress.IP,
			"uuid": ipAddress.UUID}).Debug("updated record")

		addresses = append(addresses, modelToResponse(ipAddress))
	}

	return &model.IPAddressResult{
		Node: addresses,
	}, nil
}

func (r *queryResolver) GetIPDetails(ctx context.Context, input string) (*model.IPAddressResult, error) {
	ipAddress, err := r.Repository.GetIP(input)
	if err != nil {
		r.Logger.WithField("err", err).Debug("Error while retrieving ip address")
		return &model.IPAddressResult{
			Message: "ip not found",
		}, err
	}

	responses := make([]*model.IPAddress, 0)
	modelToResponse := model.IPAddress{
		UUID:         ipAddress.UUID,
		CreatedAt:    ipAddress.CreatedAt,
		UpdatedAt:    ipAddress.UpdatedAt,
		ResponseCode: ipAddress.ResponseCode,
		IPAddress:    ipAddress.IP,
	}
	responses = append(responses, &modelToResponse)
	return &model.IPAddressResult{
		Message: "success",
		Node:    responses,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
