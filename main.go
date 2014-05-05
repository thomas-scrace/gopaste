package main

import (
    "fmt"
    "net/http"
    "github.com/thomas-scrace/gopaste/out"
)


func viewHandler(response http.ResponseWriter, request *http.Request) {
    store_dir := "/Users/tom/gopaste/store/"
    page, _ := out.GetPageForKey(store_dir, request.URL.Path)
    // handle out_err errors. was is a missing page (404) or something
    // else (500)
    _, print_err := fmt.Fprint(response, page)
    if print_err != nil {
        //log
    }
}


func main() {
    http.HandleFunc("/", viewHandler)
    http.ListenAndServe(":8080", nil)
}

