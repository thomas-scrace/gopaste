package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func makeTestDir() (string, error) {
	dir := filepath.Join(os.TempDir(), "testdir")

	mkErr := os.Mkdir(dir, 0777)
	if mkErr != nil {
		return "", mkErr
	}

	return dir, nil
}

func TestGetTextForKeyReturnsRightText(t *testing.T) {
	// arrange
	key := "testkey"
	dir, mkDirErr := makeTestDir()
	if mkDirErr != nil {
		t.Errorf(mkDirErr.Error())
	}
	defer os.RemoveAll(dir)

	filename := filepath.Join(dir, key)
	contents := []byte("Test string")
	writeErr := ioutil.WriteFile(filename, contents, 0777)
	if writeErr != nil {
		t.Errorf(writeErr.Error())
	}

	// act
	text, err := getTextForKey(dir, key)

	// assert
	if err != nil {
		t.Errorf(err.Error())
	} else if text != string(contents) {
		t.Errorf("Did not get back expected contents.")
	}
}

func TestGetTextForKeyReturnsError(t *testing.T) {
	dir, mkDirErr := makeTestDir()
	if mkDirErr != nil {
		t.Errorf(mkDirErr.Error())
	}
	defer os.RemoveAll(dir)

	text, getErr := getTextForKey(dir, "doesnotexist")
	if text != "" {
		t.Errorf(
			"Did not get back empty string from erroenous getTextForKey call")
	}
	if getErr == nil {
		t.Errorf("Did not get back error from erroneous getTextForKey call")
	}
}

func TestHash(t *testing.T) {
	data := []byte("GoPaste")
	hashed := hash(data)
	exp := "a06ee3951803be4f2be2cdb880a36b123620f87083fd8354bdc8f8fa79b17386"
	if hashed != exp {
		t.Errorf("Hash function did not return expected digest")
	}
}
