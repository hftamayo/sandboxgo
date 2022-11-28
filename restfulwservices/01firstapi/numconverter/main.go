//main.go
package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"
	"numconverter/data"
)

func main() {
	//http package has methods for dealing with requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")
		//if the request is GET with correct syntax
		if urlPathElements[1] == "roman_number" {
			number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
			if number == 0 || number > 10 {
				//if resource is not in the list, send NotFound status
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not Found"))
			} else {
				fmt.Fprintf(w, "%q", html.EscapeString(romans.Numerals[number]))
			}
		} else {
			// For all other requests, tell that client sent a bad request
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
		}
	})
	//create a server and run it on 8000 port
	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
