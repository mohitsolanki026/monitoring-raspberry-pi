package monitoring

import (
	"net/http"
	"time"
)

func GetCameraStream(ip string) (int , error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Head("http://" + ip)
	if err != nil {
		return 0,err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return 1,nil
	}

	return 0,nil
}
