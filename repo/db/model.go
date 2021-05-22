package db

import (
	"errors"
	"net"
	"strings"

	guuid "github.com/google/uuid"
	"gorm.io/gorm"
)

var ErrInvalidIPAddress = errors.New("invalid ip address specified")

type IPAddress struct {
	gorm.Model
	UUID         string `gorm:"not null"`
	ResponseCode string `gorm:"not null"`
	IPAddress    string `gorm:"index;not null"`
}

func NewIP(ipAddress string) (IPAddress, error) {
	if !isValidIPV4(ipAddress) {
		return IPAddress{}, ErrInvalidIPAddress
	}

	return IPAddress{
		IPAddress: ipAddress,
	}, nil
}

func (i IPAddress) Reverse() string {
	reversedIP := []string{}
	octets := strings.Split(i.IPAddress, ".")

	for i := range octets {
		octet := octets[len(octets)-1-i]
		reversedIP = append(reversedIP, octet)
	}
	return strings.Join(reversedIP, ".")
}

func (i *IPAddress) BeforeCreate(tx *gorm.DB) (err error) {

	id, err := guuid.NewUUID()
	if err != nil {
		return err
	}

	i.UUID = id.String()
	return
}

func isValidIPV4(ip string) bool {
	if len(ip) == 0 {
		return false
	}

	parsedIP := net.ParseIP(ip)
	if parsedIP == nil || parsedIP.To4() == nil {
		return false
	}

	return true
}
