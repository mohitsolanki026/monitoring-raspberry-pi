package monitor

import (
	"monitoringGo/monitoring"
	"monitoringGo/pkg/log"
	"monitoringGo/pkg/service"
	"os/exec"
)

func Internet() {
	download, upload, err := monitoring.GetInternetSpeed()
	errHandling(err)

	hostname, err := exec.Command("hostname").Output()
	errHandling(err)

	log.LogDataInternet("log.txt", download, upload)
	service.SendDataInternet(string(hostname),download, upload)
}
