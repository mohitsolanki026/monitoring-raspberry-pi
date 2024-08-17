package monitoring

import (
	"os/exec"
	"regexp"
	"strconv"
)

func GetCPUVoltage() (float32, error) {
	out, err := exec.Command("vcgencmd", "measure_volts").Output()
	if err != nil {
		return 0.0, err
	}

	re := regexp.MustCompile(`[0-9]+\.[0-9]+`)
	voltage := re.FindString(string(out))
	voltageFloat, err := strconv.ParseFloat(voltage, 32)
	if err != nil {
		return 0.0, err
	}
	return float32(voltageFloat), nil
}
