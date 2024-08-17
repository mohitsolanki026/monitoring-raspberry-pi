package scheduler

import (
    "log"
    "monitoringGo/monitor"

    "github.com/robfig/cron/v3"
)

func SetupCronJobs() {
    c := cron.New()

    _, err := c.AddFunc("@every 1m", func() {
        log.Println("Running every minute tasks")
        monitor.Gerneral()
    })
    if err != nil {
        log.Fatalf("Error scheduling every minute task: %v", err)
    }

    _, err = c.AddFunc("@hourly", func() {
        log.Println("Running hourly tasks")
        monitor.Internet()
    })
    if err != nil {
        log.Fatalf("Error scheduling hourly task: %v", err)
    }

    c.Start()

    select {}
}
