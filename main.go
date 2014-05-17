package main

import (
	"net/http"
	"os"
)

func main() {
	// initLogging establishes INFO and ERROR as two global writers
	// writing to stdout and stderr respectively.
	initLogging(os.Stdout, os.Stderr)

	config := getConfig()

	getHandler := getGetHandler(config.PathToStore)
	putHandler := getPutHandler(config.PathToStore)

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
