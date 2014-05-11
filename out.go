package main

import (
    "crypto/sha256"
    "encoding/hex"
    "io/ioutil"
    "path/filepath"
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

// Return data hashed using sha256 as a hex-encoded string
func hash(data []byte) string {
    digest := sha256.New()
    digest.Write(data)
    key := digest.Sum(nil)
    return hex.EncodeToString(key)
}


func SavePaste(pathToStore, text string) (string, error) {
    textBytes := []byte(text)
    key := hash(textBytes)
    path := filepath.Join(pathToStore, key)
    err := ioutil.WriteFile(path, []byte(text), 0777)
    return key, err
}
