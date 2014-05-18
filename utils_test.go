package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestConfigGetPortString(t *testing.T) {
	config := goPasteConfig{Port: 8000, PathToStore: "/home/gopaste/gopaste/"}
	if config.getPortString() != ":8000" {
		t.Errorf("getPortString method did not return expected result")
	}
}

func TestIsValidStorePathFailsOnRelativePath(t *testing.T) {
	relativePath := "this/path/is/relative/"
	expectedError := "Store path must be absolute."
	if err := isValidStorePath(relativePath); err.Error() != expectedError {
		t.Errorf(
			"isValidStorePath did not return expected error for relative path")
	}
}

func TestIsValidStorePathFailsOnNonExistentPath(t *testing.T) {
	dir, mkErr := makeTestDir(0777)
	if mkErr != nil {
		t.Errorf(mkErr.Error())
	}
	defer os.RemoveAll(dir)

	nonExistentDir := filepath.Join(dir, "doesnotexist")
	if err := isValidStorePath(nonExistentDir); err == nil {
		t.Errorf("isValidStorePath returned no error for non-existent path")
	}
}

func TestIsValidStorePathFailsOnNonWriteablePath(t *testing.T) {
	dir, mkErr := makeTestDir(0000)
	if mkErr != nil {
		t.Errorf(mkErr.Error())
	}
	defer os.RemoveAll(dir)
	if err := isValidStorePath(dir); err == nil {
		t.Errorf("isValidStorePath returned no error for non-writeable path")
	}
}

func TestIsValidStorePathReturnsNoErrorForValidPath(t *testing.T) {
	dir, mkErr := makeTestDir(0777)
	if mkErr != nil {
		t.Errorf(mkErr.Error())
	}
	defer os.RemoveAll(dir)

	if err := isValidStorePath(dir); err != nil {
		t.Errorf("isValidStorePath returns error for valid path.")
	}
}

func TestSafeUint16FromIntReturnsCorrectResultForValidInts(t *testing.T) {
	var i int
	for i = 0; i <= 65535; i++ {
		ui, err := safeUint16FromInt(i)
		if ui != uint16(i) {
			t.Errorf("safeUint16FromInt did not return correct result.")
		}
		if err != nil {
			t.Errorf("safeUint16FromInt returned error for valid input.")
		}
	}
}

func TestSafeUint16FromIntReturnsErrorForNegativeInts(t *testing.T) {
	_, err := safeUint16FromInt(-1)
	if err == nil {
		t.Errorf("safeUint16FromInt did not return error for negative int")
	}
}

func TestSafeUint16FromIntReturnsErrorForTooLargeInts(t *testing.T) {
	_, err := safeUint16FromInt(65536)
	if err == nil {
		t.Errorf("safeUint16FromInt did not return error for too large int")
	}
}