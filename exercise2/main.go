package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type urlConfig struct {
	Path string `yaml:"path" json:"path"`
	URL  string `yaml:"url" json:"url"`
}

func getFileData(filename string) []byte {
	configFileData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error when opening the config file")
		os.Exit(1)
	}

	return configFileData
}

func getFileType(filename string) string {
	lastPeriod := strings.LastIndex(filename, ".")
	return filename[lastPeriod+1:]
}

func mapHandlers(router *http.ServeMux, fileData []byte, datatype string) error {
	var routeConfigs []urlConfig
	if datatype == "yaml" {
		if err := yaml.Unmarshal(fileData, &routeConfigs); err == nil {
			return err
		}
	} else if datatype == "json" {
		if err := json.Unmarshal(fileData, &routeConfigs); err == nil {
			return err
		}
	}

	for _, config := range routeConfigs {
		router.HandleFunc(config.Path, createRedirectRouteHandler(config.URL))
	}

	return nil
}

func getRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", sayHello)

	return router
}

func main() {
	filename := flag.String("file", "routes.yaml", "a file with the initial route configuration")
	flag.Parse()

	router := getRouter()
	fileData := getFileData(*filename)

	mapHandlers(router, fileData, getFileType(*filename))
	router.HandleFunc("/routes", createAddRoutesHandler(router))

	http.ListenAndServe(":8000", router)
}
