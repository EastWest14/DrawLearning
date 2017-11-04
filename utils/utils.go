package utils

import (
	"errors"
	"io/ioutil"
	"os"
)

func loadFileToString(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", errors.New("Failed opening file: " + err.Error())
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", errors.New("Failed reading file: " + err.Error())
	}
	return string(content), nil
}
