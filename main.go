package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var pathFlag = flag.String("path", "", "specify the URL root")

func main() {
	flag.Parse()

	if *pathFlag == "" {
		log.Fatal("Missing path")
	}
	path := *pathFlag

	http.HandleFunc(fmt.Sprintf("/%s", path), func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "URL:%s", r.URL.String())
	})
	http.HandleFunc(fmt.Sprintf("/%s/health", path), func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "OK")
	})

	log.Print("Server starting at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
