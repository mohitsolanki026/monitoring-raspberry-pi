package log

import (
    "log"
    "os"
)

func InitLogger(logFile string) (*os.File, error) {
    file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        return nil, err
    }

    log.SetOutput(file)
    log.SetFlags(0)
    return file, nil
}
