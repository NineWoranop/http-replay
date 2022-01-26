package file

import (
	"fmt"
	"io/ioutil"
	"os"
)

type FileHandler struct {
	Config *HanlderConfig
}

func NewFileHandler(config *HanlderConfig) FileHandler {
	return FileHandler{Config: config}
}

func (handler *FileHandler) Read() (string, error) {
	var metricsStr, err = handler.read()

	// Retry if it has been success on previous file
	if err != nil && handler.PassedOnce() {
		// metrics.SetCounter(fh.CurrentIndex())
		metricsStr, err = handler.read()
	}
	return metricsStr, err
}

func (handler *FileHandler) read() (string, error) {
	filename := handler.Config.CurrentFilePath()
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("%s is not found", filename)
	}
	str := string(bytes)
	return str, nil
}

func (handler *FileHandler) HasNext() bool {
	if !handler.Config.CanIncrease() {
		return false
	}
	filename := handler.Config.NextFilePath()
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (handler *FileHandler) PassedOnce() bool {
	return handler.Config.PassedOnce()
}

func (handler *FileHandler) Next() {
	handler.Config.Next()
}
