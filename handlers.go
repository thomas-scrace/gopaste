// handlers.go defines functions that receive http requests and write
// http responses.
package main

import (
	"fmt"
	"net/http"
)

// getGetHandler returns a handler (closed over the path to the paste
// store) that handles requests to retrieve existing pastes.
func getGetHandler(pathToStore string) func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		key := request.URL.Path[len(pasteRoot)-1:]
		page, keyErr := getPageForKey(pathToStore, key)
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

// getPutHandler returns a handler (closed over the path to the paste
// store) that handles requests to the root. If the request is a GET
// then we should send back a paste form.  If it is a POST then we
// should take the contents of the form and save it as a paste file.
func getPutHandler(pathToStore string) func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			newPastePage := getNewPastePage()
			_, printErr := fmt.Fprint(response, newPastePage)
			if printErr != nil {
				ERROR.Println(printErr)
			}
		case "POST":
			parseErr := request.ParseForm()
			if parseErr != nil {
				internalServerError(response, parseErr)
			} else {
				text := request.FormValue("paste")
				newKey, saveErr := savePaste(pathToStore, text)
				if saveErr != nil {
					internalServerError(response, saveErr)
				} else {
					http.Redirect(
						response, request,
						pasteRoot+newKey, http.StatusFound)
				}
			}
		default:
			http.Error(
				response, "Error: Method Not Allowed (405)",
				http.StatusMethodNotAllowed)
		}
	}
}
