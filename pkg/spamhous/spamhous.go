package spamhous

import (
	"errors"
	"fmt"
	"net"

	"github.com/jorgerasillo/spamhouse/repo/db"
	"github.com/sirupsen/logrus"
)

var ErrNoSuchHost = errors.New("no such host")

func Query(ip *db.IPAddress, qChan chan *db.IPAddress, log *logrus.Logger) {
	spamhousHost := "zen.spamhous.org"
	spamIP := fmt.Sprintf("%s.%s", ip.Reverse(), spamhousHost)
	log.WithField("host", spamIP).Debug("looking up host")
	res, err := net.LookupHost(spamIP)
	if err != nil {
		log.WithField("err", err).Error("error looking up host")
		return
	}

	if len(res) > 0 {
		ip.ResponseCode = res[0]
		log.WithFields(logrus.Fields{
			"ip":       ip.IP,
			"response": ip.ResponseCode,
		}).Debug("sending ip over channel")
		qChan <- ip
	}

}
