package main

import (
	"flag"
	"net/http"
)

func createRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/intro", http.StatusTemporaryRedirect)
	})
	return router
}

func main() {
	adventureFilename := flag.String("filename", "adventure.json", "a JSON file containing the adventure story data")
	flag.Parse()

	router := createRouter()

	http.ListenAndServe(":8000", router)
}
