package log

import (
	"log"
	"time"
)

func LogData(logFile string, usbStatus, internetStatus, cameraStatus, cameraStream int, cpuTemp float64, cpuVoltage float32) {

	file, err := InitLogger(logFile)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer file.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")

	log.Printf("[%s] modem_status=%d internet_status=%d\n", timestamp, usbStatus, internetStatus)
	log.Printf("[%s] camera_status=%d camera_stream=%d\n", timestamp, cameraStatus, cameraStream)
	log.Printf("[%s] cpu_temp=%.4f cpu_voltage=%.4f\n", timestamp, cpuTemp, cpuVoltage)
	log.Println("----------------")
}

func LogDataInternet(logFile string, downloadSpeed, uploadSpeed float64) {

	file, err := InitLogger(logFile)
	if err != nil {
		LogError(logFile,"Failed to initialize logger:",err)
	}
	defer file.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")

	log.Printf("[%s] download_speed=%4f upload_speed=%4f\n", timestamp, downloadSpeed, uploadSpeed)
	log.Println("----------------")
}

func LogError(logFile string,when string, err error) {

	file, Error := InitLogger(logFile)
	if Error != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer file.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")

	log.Printf("[%s] %s error: %v\n", timestamp,when, err)
	log.Println("----------------")
}