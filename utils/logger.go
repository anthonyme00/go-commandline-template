package utils

import (
	"cltest/configs"
	"fmt"
	"log"
	"os"
	"path"
)

type Logger struct {
	logger *log.Logger
	f      *os.File
}

func newLogger(filename string) (*Logger, error) {
	outputPath := path.Join(configs.Global.PROJECT_PATH, "logs", filename+".log")
	f, err := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil,
			err
	}
	return &Logger{
		logger: log.New(f, "", log.LstdFlags),
		f:      f,
	}, nil
}

func (l *Logger) log(flag, format string, values ...interface{}) {
	message := fmt.Sprintf("[%s] %s", flag, format)
	l.logger.Println(fmt.Sprintf(message, values...))
}

func (l *Logger) LogMessage(message string) {
	l.log("INFO", "%s", message)
}

func (l *Logger) LogError(err error) {
	l.log("ERROR", "%s", err.Error())
}

func (l *Logger) LogErrorMessage(err error, message string) {
	l.log("ERRORMSG", "ERROR: %s, MESSAGE: %s", err.Error(), message)
}

func (l *Logger) LogDebug(message string) {
	l.log("DEBUG", "%s", message)
}

func (l *Logger) clean() {
	l.f.Close()
}
