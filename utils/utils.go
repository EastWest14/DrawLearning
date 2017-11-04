package utils

import (
	"errors"
	"io/ioutil"
	"os"
)

func LoadFile(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, errors.New("Failed opening file: " + err.Error())
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New("Failed reading file: " + err.Error())
	}
	return content, nil
}

func LoadFileToString(filepath string) (string, error) {
	content, err := LoadFile(filepath)
	return string(content), err
}
