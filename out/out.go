// Package out provides primitives required for outputting paste files
// stored by gopaste and rendering them as web pages.
package out

import (
    "errors"
    "io/ioutil"
    "path/filepath"
)

// getTextForKey takes:
//
// * dir, the absolute path to gopaste's file storage
// directory
// * key, a string uniquely identifying a file in dir
//
// It returns the contents of the uniquely-identified file as a byte
// array. In the case of an error the return value of the text will be
// the empty string. It is an error for the supplied directory path to
// be relative.
func getTextForKey(dir, key string) (string, error) {
    if ! filepath.IsAbs(dir) {
        return "", errors.New("The path to gopaste's store must be absolute")
    }

    path := filepath.Join(dir, key)
    text, err := ioutil.ReadFile(path)

    // text is []byte
    text_string := string(text)

    if err != nil {
        return "", err
    } else {
        return text_string, nil
    }
}


func GetPageForKey(dir, key string) (string, error) {
    text, text_err := getTextForKey(dir, key)
    if text_err != nil {
        return "", text_err
    }
    page := "<html><head><title>GoPaste</title></head><body><pre><tt>" +
            text +
            "</tt></pre></body></html>"
    return page, nil
}
