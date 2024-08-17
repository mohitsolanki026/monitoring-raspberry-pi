package monitor

import (
	"monitoringGo/config"
	"monitoringGo/monitoring"
	"monitoringGo/pkg/log"
	"monitoringGo/pkg/service"
	"os/exec"

	"strings"
)

func errHandling(err error) {
	if err != nil {
		log.LogError("log.txt", "", err)
	}
}

func Gerneral() {
	dongleName := strings.Split(config.GetEnv("DONGLE_NAMES"), " ")
	ip := config.GetEnv("IP_ADDRESS")

	stream, err := monitoring.GetCameraStream(ip)
	errHandling(err)

	voltage, err := monitoring.GetCPUVoltage()
	errHandling(err)

	dongle, err := monitoring.GetUSBStatus(dongleName)
	errHandling(err)

	temperature, err := monitoring.GetCpuTemp()
	errHandling(err)

	camera, err := monitoring.GetCameraStatus()
	errHandling(err)

	internet := monitoring.CheckInternet()

	hostname, err := exec.Command("hostname").Output()
	errHandling(err)

	log.LogData("log.txt", dongle, internet, camera, stream, temperature, voltage)
	service.SendData(string(hostname), dongle, internet, camera, stream, temperature, voltage)
}
