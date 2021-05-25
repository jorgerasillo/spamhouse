package graph

import (
	"github.com/jorgerasillo/spamhouse/graph/model"
	"github.com/jorgerasillo/spamhouse/repo/db"
)

type Status int

const (
	Failure Status = iota
	Success
	EmptyPayload
	IPNotFound
)

func (s Status) String() string {
	return []string{"Failure", "Success", "Empty", "IP not found"}[s]
}

func modelToResponse(ip db.IPAddress) *model.IPAddress {
	return &model.IPAddress{
		UUID:         ip.UUID,
		CreatedAt:    ip.CreatedAt,
		UpdatedAt:    ip.UpdatedAt,
		ResponseCode: ip.ResponseCode,
		IPAddress:    ip.IP,
	}
}
