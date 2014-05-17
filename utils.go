package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// isValidStorePath checks that the given path is (1) absolute,
// (2) extant and (3) writeable-to
func isValidStorePath(p string) error {
	if !filepath.IsAbs(p) {
		return errors.New("Store path must be absolute.")
	}

	f, openErr := os.Open(p)
	if openErr != nil {
		return openErr
	}
	f.Close()

	data := []byte{'g', 'o', 'p', 'a', 's', 't', 'e'}
	filename := "writeable-p"
	testPath := filepath.Join(p, filename)
	writeErr := ioutil.WriteFile(testPath, data, pastePerm)
	if writeErr != nil {
		return writeErr
	}
	os.Remove(testPath)

	return nil
}

type GoPasteConfig struct {
	Port        uint16 // TCP ports range 0–65535
	PathToStore string
}

func (c GoPasteConfig) getPortString() string {
	return fmt.Sprintf(":%d", c.Port)
}

func getCurrentUserHome() string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return user.HomeDir
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

func getConfig() GoPasteConfig {
	// Define the flags
	port := flag.Int("port", 8000, "The TCP port on which to serve gopaste.")
	storeDirArg := flag.String(
		"store", "$HOME/gopaste",
		"Absolute path to the directory to use for storing paste files.")

	// Parse them
	flag.Parse()

	// If storeDirArg is the default, get the user's home directory and
	// derive the correct default from it.
	var storeDir string
	if *storeDirArg == "$HOME/gopaste" {
		home := getCurrentUserHome()
		storeDir = filepath.Join(home, "gopaste")
	} else {
		storeDir = *storeDirArg
	}

	uint16Port, err := safeUint16FromInt(*port)
	if err != nil {
		log.Fatal("Specified port not in range 0–65535")
	}
	config := GoPasteConfig{Port: uint16Port, PathToStore: storeDir}

	if invalidErr := isValidStorePath(config.PathToStore); invalidErr != nil {
		log.Fatal(invalidErr)
	}
	return config
}
