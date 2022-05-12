package logger

import (
	"io"
	"log"
	"os"
)

var (
	Warning *log.Logger
	Info    *log.Logger
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("Failed to initialize logger: %v", err)
	}

	Info = createNewLogger("INFO: ", file)
	Warning = createNewLogger("WARNING: ", file)
}

func createNewLogger(prefix string, file io.Writer) *log.Logger {
	newLogger := log.New(file, prefix, log.Ldate|log.Ltime|log.Lshortfile)
	mw := io.MultiWriter(os.Stdout, file)
	newLogger.SetOutput(mw)

	return newLogger
}
