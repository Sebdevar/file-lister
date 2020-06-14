package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	cli := Cli{}
	cli.init()
	if cli.directoryLocation != "" {
		files := getFileSetFromFolder(cli.directoryLocation)
		sortFileSetBySize(files)
		err := printTable(os.Stdout, files, cli.addDirectorySizesToTotal)
		if err != nil {
			log.Fatalln("An error has occured: ", err)
		}
	}

	http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(responseWriter, request)
		return
	}
	files := getFileSetFromFolder(".")
	sortFileSetBySize(files)
	err := printTable(responseWriter, files, false)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
	}
}
