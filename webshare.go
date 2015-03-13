package main

import (
	"flag"
	"net/http"
	"log"
)

func main() {
	flag.Parse()
	dir := flag.Arg(0)
	http.Handle("/", http.FileServer(http.Dir(dir)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}