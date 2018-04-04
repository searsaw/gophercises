package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func createRedirectRouteHandler(redirectURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)
	}
}

func createAddRoutesHandler(router *http.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatalf("Error when reading the response body in the YAML handler: %s\n", err.Error())
			return
		}

		var datatype string
		switch r.Header.Get("Content-Type") {
		case "application/x-yaml":
			datatype = "yaml"
		case "application/json":
			datatype = "json"
		default:
			datatype = "yaml"
		}

		err = mapHandlers(router, requestBody, datatype)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Malformed YAML"))
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
