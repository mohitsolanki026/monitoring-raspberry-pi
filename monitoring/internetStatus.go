package monitoring

import (
	"net"
	"time"
)

func CheckInternet() int {
	sites := []string{"google.com", "cloudflare.com", "github.com"}
	sleepDelay := 10 * time.Second
	for _, site := range sites {
		if isSiteReachable(site) {
			return 1 
		}
		time.Sleep(sleepDelay)
	}
	return 0 
}

func isSiteReachable(site string) bool {
	timeout := 3 * time.Second
	conn, err := net.DialTimeout("tcp", site+":80", timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
