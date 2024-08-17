package monitoring

import (
	"os/exec"
	"strconv"
	"strings"
)

func GetCpuTemp() (float64, error) {
	out, err := exec.Command("cat", "/sys/class/thermal/thermal_zone0/temp").Output()
	if err != nil {
		return 0, err
	}
	temp := strings.TrimSpace(string(out))
	// return temp[:len(temp)-3] + "." + temp[len(temp)-3:], nil

	tempFloat, err := strconv.ParseFloat(temp, 32)
	if err != nil {
		return 0, err
	}
	return float64(tempFloat) / 1000, nil
}
