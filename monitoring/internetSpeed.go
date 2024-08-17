package monitoring

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type SpeedTestResult struct {
	Download float64 `json:"download"`
	Upload   float64 `json:"upload"`
}

// Function to get internet speed using speedtest-cli
func GetInternetSpeed() (float64, float64, error) {
	// Run the speedtest-cli command with --json flag
	cmd := exec.Command("speedtest-cli", "--json")
	output, err := cmd.Output()
	if err != nil {
		return 0, 0, fmt.Errorf("failed to run speedtest-cli: %v", err)
	}

	// Parse the JSON output
	var result SpeedTestResult
	if err := json.Unmarshal(output, &result); err != nil {
		return 0, 0, fmt.Errorf("failed to parse JSON: %v", err)
	}

	// Convert speeds from bps to Mbps
	downloadSpeedMbps := result.Download / 1000000
	uploadSpeedMbps := result.Upload / 1000000

	return downloadSpeedMbps, uploadSpeedMbps, nil
}
