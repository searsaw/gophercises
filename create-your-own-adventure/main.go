package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/searsaw/gophercises/create-your-own-adventure/adventure"
)

func createRouter() *http.ServeMux {
	router := http.NewServeMux()
	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.Redirect(w, r, "/intro", http.StatusTemporaryRedirect)
	// })
	return router
}

func main() {
	adventureFilename := flag.String("adventure", "adventure.json", "a JSON file containing the adventure story data")
	templateFilename := flag.String("template", "template.html", "a Golang template file used to show each page")
	port := flag.Int("port", 8000, "port the server runs on")
	flag.Parse()

	fileData, err := ioutil.ReadFile(*adventureFilename)
	if err != nil {
		fmt.Printf("There was a problem when opening the story file: %s\n", err.Error())
		os.Exit(1)
	}

	router := createRouter()
	adventureMux, err := adventure.NewHttpAdventure(fileData, *templateFilename)
	if err != nil {
		fmt.Printf("There was a problem when setting up the adventure mux: %s\n", err.Error())
		os.Exit(1)
	}
	router.Handle("/", adventureMux)

	fmt.Printf("The server is now running on port %d.\n", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), router)
}
