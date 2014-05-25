package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
    "syscall"
	"testing"
)

func makeTestDir(mode uint32) (string, error) {
	fileMode := os.FileMode(mode)

	dirPath := filepath.Join(os.TempDir(), "testdir")

	mkErr := os.Mkdir(dirPath, fileMode)
	if mkErr != nil {
		return "", mkErr
	}

	return dirPath, nil
}

func TestGetTextForKeyReturnsRightText(t *testing.T) {
	// arrange
	key := "testkey"
	dir, mkDirErr := makeTestDir(0777)
	if mkDirErr != nil {
		t.Errorf(mkDirErr.Error())
	}
	defer os.RemoveAll(dir)

    // Change to the new directory, remembering to change
    // back at the end.
    pwd, wdErr := os.Getwd()
    if wdErr != nil {
        t.Errorf(wdErr.Error())
    }
    syscall.Chdir(dir)
    defer syscall.Chdir(pwd)

	filename := filepath.Join(dir, key)
	contents := []byte("Test string")
	writeErr := ioutil.WriteFile(filename, contents, 0777)
	if writeErr != nil {
		t.Errorf(writeErr.Error())
	}

	text, err := getTextForKey(key)

	if err != nil {
		t.Errorf(err.Error())
	} else if text != string(contents) {
		t.Errorf("Did not get back expected contents.")
	}
}

func TestGetTextForKeyReturnsError(t *testing.T) {
	dir, mkDirErr := makeTestDir(0777)
	if mkDirErr != nil {
		t.Errorf(mkDirErr.Error())
	}
	defer os.RemoveAll(dir)

	text, getErr := getTextForKey("doesnotexist")
	if text != "" {
		t.Errorf(
			"Did not get back empty string from erroenous getTextForKey call")
	}
	if getErr == nil {
		t.Errorf("Did not get back error from erroneous getTextForKey call")
	}
}

func TestSavePasteCreatesRightKey(t *testing.T) {
	dir, mkDirErr := makeTestDir(0777)
	if mkDirErr != nil {
		t.Errorf(mkDirErr.Error())
	}
	defer os.RemoveAll(dir)

	testText := []byte("GoPaste")
	key, saveErr := savePaste(string(testText))
	if saveErr != nil {
		t.Errorf(saveErr.Error())
	}
	if key != hash(testText) {
		t.Errorf("savePaste did not correctly create key.")
	}
}

func TestSavePasteSavesPaste(t *testing.T) {
    // Set up a temporary test dir
	dir, mkDirErr := makeTestDir(0777)
	if mkDirErr != nil {
		t.Errorf(mkDirErr.Error())
	}
	defer os.RemoveAll(dir)

    // Change to the new directory, remembering to change
    // back at the end.
    pwd, wdErr := os.Getwd()
    if wdErr != nil {
        t.Errorf(wdErr.Error())
    }
    syscall.Chdir(dir)
    defer syscall.Chdir(pwd)

    // Save some text into a file
	testText := "GoPaste"
	key, saveErr := savePaste(testText)
	if saveErr != nil {
		t.Errorf(saveErr.Error())
	}

    // Make sure we save what we thought we should save
	text, readErr := ioutil.ReadFile(key)
	if readErr != nil {
		t.Errorf(readErr.Error())
	}
	if string(text) != testText {
		t.Errorf("Did not find expected text in saved paste")
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
