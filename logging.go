// logging.go establishes the global logging writers used everywhere.
package main

import (
	"io"
	"log"
)

var (
	INFO  *log.Logger
	ERROR *log.Logger
)

func initLogging(infoHandle io.Writer, errorHandle io.Writer) {
	INFO = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	ERROR = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
