// utils.go contains various utility functions and types.
package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// die is a wrapper around error-returning functions that exits
// if the error is non-nil and logs the error message.
func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Check that p is writeable-to.
func testPwdWriteable() error {
	data := []byte("gopaste")
	filename := "writeable-p"

	writeErr := ioutil.WriteFile(filename, data, pastePerm)
	if writeErr != nil {
		return writeErr
	}
	os.Remove(filename)

	return nil
}

type goPasteConfig struct {
	Port        uint16 // TCP ports range 0–65535
	PathToStore string
}

func (c goPasteConfig) getPortString() string {
	return fmt.Sprintf(":%d", c.Port)
}

// Cast i to a uint16 without wrapping. If the number is unrepresentable
// as a uint16 an error is returned.
func safeUint16FromInt(i int) (uint16, error) {
	if 0 <= i && i <= 65535 {
		return uint16(i), nil
	} else {
		return 0, errors.New("'%d' is out of range 0–65535.")
	}
}

// Return a copy of s safe for use in HTML. This is done by replacing
// charaters &, >, <, ', " with their HTML escape sequences.
func escape(s string) string {
	replacer := strings.NewReplacer(
		"&", "&amp;",
		">", "&gt;",
		"<", "&lt;",
		"'", "&#39;",
		`"`, "&#34;")
	return replacer.Replace(s)
}

// Response with a 500 and log the error.
func internalServerError(response http.ResponseWriter, e error) {
	ERROR.Println(e)
	http.Error(
		response,
		"Error: Internal Server Error (500)",
		http.StatusInternalServerError)
}

func getConfig() goPasteConfig {
	// Define the flags
	port := flag.Int("port", 8000, "The TCP port on which to serve gopaste.")
	storeDirArg := flag.String(
		"store", "REQUIRED",
		"Absolute path to the directory to use for storing paste files.")

	// Parse them
	flag.Parse()

	var storeDir string
	if *storeDirArg == "REQUIRED" {
		log.Fatal(
			"You must specify the directory to use as the store using --store.")
	} else {
		storeDir = *storeDirArg
	}

	uint16Port, err := safeUint16FromInt(*port)
	if err != nil {
		log.Fatal("Specified port not in range 0–65535")
	}

	if !filepath.IsAbs(storeDir) {
		log.Fatal("Store path must be absolute.")
	}

	config := goPasteConfig{Port: uint16Port, PathToStore: storeDir}

	return config
}
