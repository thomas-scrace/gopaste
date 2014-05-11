package main

import (
	"fmt"
	"net/http"
	"os"
)

// URLs for stored pastes are rooted here
var pasteRoot string = "/paste/"


func respondWithForm(response http.ResponseWriter) {
    page := "<!DOCTYPE html>\n" +
            "<head>\n\t" +
            "<meta charset=\"UTF-8\">\n\t" +
            "<title>GoPaste</title>\n" +
            "</head>\n\n" +
            "<body>\n\t" +
            "<form action=\"/\" method=\"post\">" +
            "<input type=\"textarea\" name=\"paste\">" +
            "<input type=\"submit\" value=\"Save\">" +
            "</form>" +
            "</body>"

    _, printErr := fmt.Fprint(response, page)
    if printErr != nil {
        ERROR.Println(printErr)
    }
}

func internalServerError(response http.ResponseWriter, e error) {
    ERROR.Println(e)
    http.Error(
        response,
        "Error: Internal Server Error (500)",
        http.StatusInternalServerError)
}


func makeNewPaste(response http.ResponseWriter, request *http.Request, pathToStore string) {
    parseErr := request.ParseForm()
    if parseErr != nil {
        internalServerError(response, parseErr)
    } else {
        text := request.FormValue("paste")
        newKey, saveErr := SavePaste(pathToStore, text)
        if saveErr != nil {
            internalServerError(response, saveErr)
        } else {
            http.Redirect(
                response, request,
                pasteRoot + newKey, http.StatusFound)
        }
    }
}


func main() {
	initLogging(os.Stdout, os.Stderr)

	config := getConfig()
    pathToStore := string(config.PathToStore)

	// We define a function to handle get requests that closes over pathToStore
	getHandler := func(response http.ResponseWriter, request *http.Request) {
        key := request.URL.Path[len(pasteRoot) - 1:]
		page, keyErr := GetPageForKey(pathToStore, key)
        if keyErr != nil {
            ERROR.Println(keyErr)
            http.NotFound(response, request)
        } else {
            _, printErr := fmt.Fprint(response, page)
            if printErr != nil {
                ERROR.Println(printErr)
            }
        }
	}

    // putHandler handles requests to the root. If the request is a GET
    // then we should send back a paste form. If it is a POST then we should
    // take the contents of the form and save it as a paste file.
    // Again, we define it as an anonymous function so that we can close
    // over pathToStore.
    putHandler := func(response http.ResponseWriter, request *http.Request) {
        switch request.Method {
        case "GET":
            respondWithForm(response)
        case "POST":
            makeNewPaste(response, request, pathToStore)
        default:
            http.Error(
                response, "Error: Method Not Allowed (405)",
                http.StatusMethodNotAllowed)
        }
    }

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
