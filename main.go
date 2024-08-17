package main

import (
	"fmt"
	"monitoringGo/config"
	"monitoringGo/monitor"
	"monitoringGo/scheduler"
)

func main() {
	fmt.Println("server started ")
	config.LoadEnv()
	monitor.Gerneral()
	scheduler.SetupCronJobs()
}
