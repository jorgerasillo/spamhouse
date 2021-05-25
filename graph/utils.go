package graph

import (
	"github.com/jorgerasillo/spamhouse/graph/model"
	"github.com/jorgerasillo/spamhouse/repo/db"
)

type Status int

const (
	Failure Status = iota
	Success
)

func (s Status) String() string {
	return []string{"Failure", "Success"}[s]
}

func modelToResponse(ip db.IPAddress) *model.IPAddress {
	return &model.IPAddress{
		ResponseCode: ip.ResponseCode,
		IPAddress:    ip.IP,
		UUID:         ip.UUID,
	}
}
