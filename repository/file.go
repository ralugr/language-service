package repository

import (
	"github.com/ralugr/language-service/logger"
	"log"
	"os"
)

type File struct {
	file *os.File
}

func New(filename string) File {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}
	logger.Warning.Printf("Initialized repository with file %v", f.Name())
	return File{f}
}
func (f File) Read() ([]byte, error) {
	data, err := os.ReadFile(f.file.Name())

	if err != nil {
		logger.Warning.Printf("Could not read from file %v", f.file.Name())
	} else {
		logger.Info.Printf("Read %v bytes from file %v", string(data), f.file.Name())
	}

	return data, err
}

func (f File) Write(bytes []byte, overwrite bool) error {
	if overwrite {
		if err := f.file.Truncate(0); err != nil {
			logger.Warning.Printf("Could not truncate file %v, error %v", f.file.Name(), err)
		}
		if _, err := f.file.Seek(0, 0); err != nil {
			logger.Warning.Printf("Could not reset the offset of the file %v, error %v", f.file.Name(), err)
		}
	}

	if _, err := f.file.Write(bytes); err != nil {
		logger.Warning.Printf("Could not write bytes %v to file %v, error%v", string(bytes), f.file.Name(), err)
	}

	logger.Warning.Printf("Wrote %v in file %v", string(bytes), f.file.Name())
	return nil
}

func (f File) Close() {
	if err := f.file.Close(); err != nil {
		logger.Warning.Printf("Could not close file %v, error%v", f.file.Name(), err)
	}
	logger.Info.Printf("Closing file %v", f.file.Name())
}
