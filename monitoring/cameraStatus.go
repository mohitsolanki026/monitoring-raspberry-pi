package monitoring

import (
    "os/exec"
    "strings"
)

func GetCameraStatus() (int, error) {
    out, err := exec.Command("ip", "a").Output()
    if err != nil {
        return 0, err
    }
    
    if strings.Contains(string(out), "169.254") {
        return 1, nil
    }
    return 0, nil
}
