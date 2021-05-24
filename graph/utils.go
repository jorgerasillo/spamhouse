package graph

import (
	"github.com/jorgerasillo/spamhouse/graph/model"
	"github.com/jorgerasillo/spamhouse/repo/db"
)

func modelToResponse(ip db.IPAddress) *model.IPAddress {
	return &model.IPAddress{
		ResponseCode: ip.ResponseCode,
		IPAddress:    ip.IP,
		UUID:         ip.UUID,
	}
}
