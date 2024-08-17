package service

import (
	"fmt"
	"monitoringGo/config"
	"monitoringGo/pkg/log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func SendData(hostname string, modemStatus, internetStatus, cameraStatus, cameraStream int, cpuTemp float64, cpuVoltage float32) {
	urlStr := config.GetEnv("API_URL")

	baseURL, err := url.Parse(urlStr + strings.TrimSuffix(hostname, "\r\n"))
	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
		return
	}

	params := url.Values{}
	params.Add("camera_status", strconv.Itoa(cameraStatus))
	params.Add("camera_stream", strconv.Itoa(cameraStream))
	params.Add("modem_status", strconv.Itoa(modemStatus))
	params.Add("cpu_temp", strconv.Itoa(int(cpuTemp)))
	params.Add("cpu_voltage", strconv.Itoa(int(cpuVoltage)))
	params.Add("internet_status", strconv.Itoa(internetStatus))

	baseURL.RawQuery = params.Encode()

	resp, err := http.Get(baseURL.String())
	if err != nil {
		log.LogError("log.txt", "send data", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.LogError("log.txt", "response data", fmt.Errorf("%d", resp.StatusCode))
	}
}

func SendDataInternet(hostname string, downloadSpeed, uploadSpeed float64) {
	urlStr := config.GetEnv("API_URL")

	baseURL, err := url.Parse(urlStr + strings.TrimSuffix(hostname, "\r\n"))
	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
		return
	}

	params := url.Values{}
	params.Add("download_speed", strconv.Itoa(int(downloadSpeed)))
	params.Add("upload_speed", strconv.Itoa(int(uploadSpeed)))

	baseURL.RawQuery = params.Encode()

	resp, err := http.Get(baseURL.String())
	if err != nil {
		log.LogError("log.txt", "send data", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.LogError("log.txt", "response data", fmt.Errorf("%d", resp.StatusCode))
	}
	fmt.Println(http.StatusOK)
}
