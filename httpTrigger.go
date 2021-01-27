// Package helloworld provides a set of Cloud Functions samples.
package helloworld

import (
        "encoding/json"
        "fmt"
        "html"
        "net/http"
        "io/ioutil"
        "log"
)

// HelloHTTP is an HTTP Cloud Function with a request parameter.
// function - 1
func HelloHTTP(w http.ResponseWriter, r *http.Request) {
        var d struct {
                Name string `json:"name"`
        }
        if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
                fmt.Fprint(w, "Hello, World!")
                return
        }
        if d.Name == "" {
                fmt.Fprint(w, "Hello, World!")
                return
        }
        fmt.Fprintf(w, "Hello, %s!", html.EscapeString(d.Name))
}

//function - 2
func HelloContentType(w http.ResponseWriter, r *http.Request) {
        var name string

        switch r.Header.Get("Content-Type") {
        case "application/json":
                var d struct {
                        Name string `json:"name"`
                }
                err := json.NewDecoder(r.Body).Decode(&d)
                if err != nil {
                        log.Printf("error parsing application/json: %v", err)
                } else {
                        name = d.Name
                }
        case "application/octet-stream":
                body, err := ioutil.ReadAll(r.Body)
                if err != nil {
                        log.Printf("error parsing application/octet-stream: %v", err)
                } else {
                        name = string(body)
                }
        case "text/plain":
                body, err := ioutil.ReadAll(r.Body)
                if err != nil {
                        log.Printf("error parsing text/plain: %v", err)
                } else {
                        name = string(body)
                }
        case "application/x-www-form-urlencoded":
                if err := r.ParseForm(); err != nil {
                        log.Printf("error parsing application/x-www-form-urlencoded: %v", err)
                } else {
                        name = r.FormValue("name")
                }
        }

        if name == "" {
                name = "World"
        }

        fmt.Fprintf(w, "Hello, %s!", html.EscapeString(name))
}