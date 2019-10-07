package log

import (
	"log"
	"os"
)

func Logger() *log.Logger {
	logger := log.New(os.Stdout, "DEBUG", log.Ldate|log.Ltime)
	return logger
}
