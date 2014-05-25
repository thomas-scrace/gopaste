package main

import (
	"os"
	"syscall"
	"testing"
)

func TestConfigGetPortString(t *testing.T) {
	config := goPasteConfig{Port: 8000, PathToStore: "/home/gopaste/gopaste/"}
	if config.getPortString() != ":8000" {
		t.Errorf("getPortString method did not return expected result")
	}
}

func testTestPwdWriteableFailsOnNonWriteablePwd(t *testing.T) {
	// Make a test directory that is non-writeable
	dir, mkErr := makeTestDir(0000)
	if mkErr != nil {
		t.Errorf(mkErr.Error())
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

	// Ask testPwdWriteable if dir is writeable
	if err := testPwdWriteable(); err == nil {
		t.Errorf("testPwdWriteable returned no error for non-writeable path")
	}
}

func TestTestPwdWriteableReturnsNoErrorForValidPath(t *testing.T) {
	dir, mkErr := makeTestDir(0777)
	if mkErr != nil {
		t.Errorf(mkErr.Error())
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

	if err := testPwdWriteable(); err != nil {
		t.Errorf("testPwdWriteable returns error for valid path.")
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
