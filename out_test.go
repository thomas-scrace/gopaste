package main

import (
    "io/ioutil"
    "os"
    "path/filepath"
    "testing"
)


func TestGetTextForKeyReturnsRightText(t *testing.T) {
    // arrange
    key := "testkey"
    dir := filepath.Join(os.TempDir(), "testdir")

    mk_err := os.Mkdir(dir, 0777)
    if mk_err != nil {
        t.Errorf(mk_err.Error())
    }

    filename := filepath.Join(dir, key)
    contents := []byte("Test string")
    write_err := ioutil.WriteFile(filename, contents, 0777)
    if write_err != nil {
        t.Errorf(write_err.Error())
    }

    // act
    text, err := getTextForKey(dir, key)

    // assert
    if err != nil {
        t.Errorf(err.Error())
    } else if text != string(contents) {
        t.Errorf("Did not get back expected contents.")
    }

    // cleanup
    os.RemoveAll(dir)
}
