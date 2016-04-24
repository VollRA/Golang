package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	hah, err := ioutil.ReadAll(r.Body)
	if err != nil {

		fmt.Fprintf(w, "%s", err)

	}
	fmt.Fprintf(w, "I read string - %s!", hah)
}

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/read", readHandler)

	s := &http.Server{
		Addr:           ":8081",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())

}
