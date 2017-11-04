package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadFileToString(t *testing.T) {
	const MISS_PATH = "./noexistent"
	_, err := os.Open(MISS_PATH)
	if err == nil {
		t.Errorf("File at path %s exists, but shouldn't for the purpose of testing. Please remove it.", MISS_PATH)
		return
	}
	content, err := loadFileToString(MISS_PATH)
	if err == nil {
		t.Errorf("loadFileToString doesn't return an error while trying to open a nonexistent file at path: %s", MISS_PATH)
	}
	if content != "" {
		t.Errorf("Calling loadFileToString on a non-existent filepath should return empty string. Instead got: %s", content)
	}

	fContent := []byte("You broke my heart, Fredo.")
	tDir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Errorf("Failed creating temp directory: %s", err.Error())
	}
	defer os.RemoveAll(tDir)

	tFile := filepath.Join(tDir, "test_file_to_load")
	err = ioutil.WriteFile(tFile, fContent, 0200)
	if err != nil {
		t.Errorf("Failed to write temporary file: %s", err.Error())
	}
	content, err = loadFileToString(tFile)
	if err == nil {
		t.Error("Expected permission error loading file to string. Got no error.")
	}
	if content != "" {
		t.Errorf("Expected file content to be returned as: %s because of permission error, got: %s", "", content)
	}

	tFile2 := filepath.Join(tDir, "test_file_to_load2")
	err = ioutil.WriteFile(tFile2, fContent, 0666)
	if err != nil {
		t.Errorf("Failed to write temporary file: %s", err.Error())
	}
	content, err = loadFileToString(tFile2)
	if err != nil {
		t.Errorf("Failed loading file to string: %s", err.Error())
	}
	if content != string(fContent) {
		t.Errorf("Expected file content loaded to be: %s, got: %s", string(fContent), content)
	}
}
