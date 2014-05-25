// handlers.go defines functions that receive http requests and write
// http responses.
package main

import (
	"fmt"
	"net/http"
)

// getHandler handles requests to retrive existing pastes.
func getHandler(response http.ResponseWriter, request *http.Request) {
		key := request.URL.Path[len(pasteRoot)-1:]
		page, keyErr := getPageForKey(key)
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

// putHandler handles requests made to the root path. It the request
// is a GET then we return a page containing a form. If it is a POST
// then we save the new paste.
func putHandler(response http.ResponseWriter, request *http.Request) {
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
				newKey, saveErr := savePaste(text)
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
