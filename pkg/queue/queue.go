package queue

import (
	"fmt"

	"github.com/jorgerasillo/spamhouse/repo"
)

type Queue struct {
	qChan chan string
	repo  repo.Repository
}

func New(qChan chan string, repo repo.Repository) *Queue {
	return &Queue{
		qChan: qChan,
	}
}

func (q *Queue) Process() {
	for {
		ipAddress := <-q.qChan
		fmt.Printf(ipAddress)

		// get ip from spamhaus

		// query if it exists

		// save if it does

		// create if it doesn't
	}
}
