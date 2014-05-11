package main

import (
    "errors"
    "flag"
    "fmt"
    "log"
    "os/user"
    "path/filepath"
    "strings"
)

type AbsPath string

func (p AbsPath) isValid() bool {
    return filepath.IsAbs(string(p))
}


type GoPasteConfig struct {
    Port uint16 // TCP ports range 0–65535
    PathToStore AbsPath
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


func safeUint16FromInt(i int) (uint16, error){
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


func getConfig() GoPasteConfig {
    // Define the flags
    port := flag.Int("port", 80, "The TCP port on which to serve gopaste.")
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
    config := GoPasteConfig{Port:uint16Port, PathToStore:AbsPath(storeDir)}

    if ! config.PathToStore.isValid() {
        log.Fatalf("\"%d\" is invalid as an absolute file path.")
    }
    return config
}
