package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	fileRoot := flag.String("root", ".", "file system root")
	bind := flag.String("http", ":8080", "http binding")
	flag.Parse()

	err := os.Chdir(*fileRoot)
	if err != nil {
		log.Fatal(err)
	}

	handler := http.FileServer(http.Dir(*fileRoot))
	http.ListenAndServe(*bind, handler)
}
