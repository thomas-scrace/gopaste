package main

import (
    "io/ioutil"
    "path/filepath"
    "crypto/sha512"
)


// getTextForKey takes:
//
// * pathToStore, the absolute path to gopaste's file storage
// directory
// * key, a string uniquely identifying a file in dir
//
// It returns the contents of the uniquely-identified file as an
// escaped-for-html string. In the case of an error the return value of
// the text will be the empty string.
func getTextForKey(pathToStore, key string) (string, error) {
    path := filepath.Join(pathToStore, key)
    text, err := ioutil.ReadFile(path)

    if err != nil {
        return "", err
    }

    // text is []byte
    textString := string(text)
    escapedTextString := escape(textString)

    return escapedTextString, nil
}


func GetPageForKey(pathToStore, key string) (string, error) {
    text, textErr := getTextForKey(pathToStore, key)
    if textErr != nil {
        return "", textErr
    }

    page := "<!DOCTYPE html>\n" +
            "<head>\n\t" +
            "<meta charset=\"UTF-8\">\n\t" +
            "<title>GoPaste</title>\n" +
            "</head>\n\n" +
            "<body>\n\t" +
            "<pre><tt>" +
            text +
            "\t</tt></pre>\n" +
            "</body>"

    return page, nil
}


func SavePaste(pathToStore, text string) (string, error) {
    textBytes := []byte(text)
    digest := sha512.Sum512(textBytes)
    // digest is a fixed-length array, so first we have to convert it
    // to a slice, and then to a string
    key := string(digest[:])
    path := filepath.Join(pathToStore, key)
    err := ioutil.WriteFile(path, []byte(text), 0777)
    return key, err
}
