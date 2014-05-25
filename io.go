// io.go contains functions directly concerned with saving and reading
// pasted text to and from the filesystem. Nothing in here knows
// anything about html or http.
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"os"
)

// getTextForKey opens the paste file in pwd with the same name
// as the key argument and returns its contents as a string.
// If the file could not be read for any reason, the
// return values will be the empty string and whatever error was
// encountered during the attempt to read the file.
func getTextForKey(key string) (string, error) {
	text, err := ioutil.ReadFile(key)

	if err != nil {
		return "", err
	}

	// text is []byte
	textString := string(text)
	return textString, nil
}

// hash returns data hashed using sha256 as a hex-encoded string
func hash(data []byte) string {
	digest := sha256.New()
	digest.Write(data)
	key := digest.Sum(nil)
	return hex.EncodeToString(key)
}

// savePaste saves text to a file in pwd whose filename
// is a hash of the text. The filename is returned along with any error
// encountered while trying to save.
func savePaste(text string) (string, error) {
	textBytes := []byte(text)
	key := hash(textBytes)

	// We try to open the file to see if it already exists.
	var err error
	file, openErr := os.Open(key)
	if openErr != nil {
		err = ioutil.WriteFile(key, []byte(text), pastePerm)
	} else {
		// No need to write anything; we succcessfully opened the file.
		file.Close()
		err = nil
	}

	return key, err
}
