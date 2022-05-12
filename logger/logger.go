package logger

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var (
	Warning    *log.Logger
	Info       *log.Logger
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func init() {
	file, err := os.OpenFile(basepath+"/../logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("Failed to initialize logger: %v", err)
	}

	Info = createNewLogger("INFO: ", file)
	Warning = createNewLogger("WARNING: ", file)
}

func createNewLogger(prefix string, file io.Writer) *log.Logger {
	newLogger := log.New(file, "LANGUAGE: "+prefix, log.Ldate|log.Ltime|log.Lshortfile)
	mw := io.MultiWriter(os.Stdout, file)
	newLogger.SetOutput(mw)

	return newLogger
}
