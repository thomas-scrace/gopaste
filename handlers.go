package main

import (
	"fmt"
	"net/http"
)

// Return a handler closed over a pathToStore to handle requests for
// existing pastes.
func getGetHandler(pathToStore string) func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		key := request.URL.Path[len(pasteRoot)-1:]
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
}

// Return a handler closed over a pathToStore to handles requests to the
// root. If the request is a GET then we should send back a paste form.
// If it is a POST then we should take the contents of the form and save
// it as a paste file.
func getPutHandler(pathToStore string) func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
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
				pasteRoot+newKey, http.StatusFound)
		}
	}
}

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
