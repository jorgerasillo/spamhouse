package spamhous

import (
	"errors"
	"fmt"
	"net"

	"github.com/jorgerasillo/spamhouse/repo/db"
)

var ErrNoSuchHost = errors.New("no such host")

func Query(ip *db.IPAddress, qChan chan *db.IPAddress) {
	spamhousHost := "zen.spamhous.org"
	spamIP := fmt.Sprintf("%s.%s", ip.Reverse(), spamhousHost)
	fmt.Printf("looking up host %s\n", spamIP)
	res, err := net.LookupHost(spamIP)
	if err != nil {
		fmt.Printf("error looking up host: %v, err. Err: %v", err, ErrNoSuchHost)
		return
	}

	if len(res) > 0 {
		ip.ResponseCode = res[0]
		fmt.Printf("sending ip over channel: %s with response code: %s\n", ip.IP, ip.ResponseCode)
		qChan <- ip
	}

}
