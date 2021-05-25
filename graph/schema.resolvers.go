package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/jorgerasillo/spamhouse/graph/generated"
	"github.com/jorgerasillo/spamhouse/graph/model"
	"github.com/jorgerasillo/spamhouse/pkg/spamhous"
	"github.com/jorgerasillo/spamhouse/repo/db"
	"github.com/sirupsen/logrus"
)

func (r *mutationResolver) Enqueue(ctx context.Context, input []string) (*model.Result, error) {
	errors := make([]*string, 0)
	for _, i := range input {
		ip, err := db.NewIP(i)
		r.Logger.WithFields(logrus.Fields{"ip": i, "uuid": ip.UUID}).Debug("processing input")

		if err != nil {
			r.Logger.WithFields(logrus.Fields{
				"err": err,
				"ip":  ip,
			}).Error("unable to create new ip address")
			e := err.Error()
			errors = append(errors, &e)
		}
		// query spamhous and send result to channel
		// processor will pick up changes from the channel and
		// save in db
		go spamhous.Query(&ip, r.QChan)
	}

	result := model.Result{}
	result.Status = Success.String()
	if len(errors) > 0 {
		result.Status = Failure.String()
	}

	result.Errors = errors
	return &result, nil
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
