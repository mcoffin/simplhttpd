package main

import (
	"flag"
	"net/http"
)

func main() {
	fileRoot := flag.String("root", ".", "file system root")
	bind := flag.String("http", ":8080", "http binding")
	flag.Parse()

	handler := http.FileServer(http.Dir(*fileRoot))
	http.ListenAndServe(*bind, handler)
}
