package queue

import (
	"errors"

	"github.com/jorgerasillo/spamhouse/repo"
	"github.com/jorgerasillo/spamhouse/repo/db"
	"github.com/sirupsen/logrus"
)

type Queue struct {
	repo   repo.Repository
	qChan  chan *db.IPAddress
	logger *logrus.Logger
}

func New(repo repo.Repository, qChan chan *db.IPAddress, log *logrus.Logger) Queue {
	return Queue{
		repo:   repo,
		qChan:  qChan,
		logger: log,
	}
}
func (q Queue) Process() {
	q.logger.Debug("Starting queue for processing")
	for {
		ip := <-q.qChan
		if ip == nil {
			q.logger.Error("unable to process nil id")
			continue
		}

		logEntry := q.logger.WithFields(logrus.Fields{"ip": ip.IP, "uuid": ip.UUID})

		// query if it exists
		oldIP, err := q.repo.GetIP(ip.String())
		if err != nil {
			// check if not found
			if errors.Is(err, db.ErrRecordNotFound) {
				logEntry.Debug("record was not found, creating new entry")
				ipAddress, err := q.repo.AddIP(*ip)
				if err != nil {
					q.logger.WithFields(logrus.Fields{
						"err": err,
						"ip":  ip,
					}).Error("unable to add new ip address")
					continue
				}
				q.logger.WithFields(logrus.Fields{
					"ip":   ipAddress.IP,
					"uuid": ipAddress.UUID}).Debug("new record added")

			}
			logEntry.Debug("db error attempting to look for record")
			continue

		}

		oldIP.ResponseCode = ip.ResponseCode
		q.logger.WithFields(logrus.Fields{
			"ip":   oldIP.IP,
			"uuid": oldIP.UUID}).Debug("record existed, updating entry")

		// save if it does
		ipAddress, err := q.repo.UpdateIP(oldIP)
		if err != nil {
			q.logger.WithFields(logrus.Fields{
				"err": err,
				"ip":  ip,
			}).Error("unable to update ip address")
			continue
		}
		q.logger.WithFields(logrus.Fields{
			"ip":   ipAddress.IP,
			"uuid": ipAddress.UUID}).Debug("updated record")

	}
}
