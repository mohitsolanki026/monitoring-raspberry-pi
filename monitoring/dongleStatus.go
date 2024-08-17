package monitoring

import (
	"os/exec"
	"strings"
)

func GetUSBStatus(dongleNames []string) (int, error) {
	out, err := exec.Command("lsusb").Output()
	if err != nil {
		return 0, err
	}

	for _, dongleName := range dongleNames {
		if strings.Contains(strings.ToLower(string(out)), strings.ToLower(dongleName)) {
			return 1, nil
		}
	}

	return 0, nil
}