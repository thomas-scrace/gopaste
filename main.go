package main

import (
	"net/http"
	"os"
)


func main() {
	initLogging(os.Stdout, os.Stderr)

	config := getConfig()
    pathToStore := string(config.PathToStore)

    getHandler := getGetHandler(pathToStore)
    putHandler := getPutHandler(pathToStore)

	http.HandleFunc("/", putHandler)
	http.HandleFunc(pasteRoot, getHandler)

	INFO.Printf(
		"Starting server on port %d serving from %q.",
		config.Port, config.PathToStore)

	httpErr := http.ListenAndServe(config.getPortString(), nil)
	if httpErr != nil {
		ERROR.Println(httpErr)
	}
}
