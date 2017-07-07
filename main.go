package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var pathFlag = flag.String("path", "", "specify the URL root")

func main() {
	flag.Parse()

	if *pathFlag == "" {
		log.Fatal("Missing path")
	}
	path := *pathFlag

	http.HandleFunc(fmt.Sprintf("/%s", path), func(rw http.ResponseWriter, r *http.Request) {
		host, err := os.Hostname()
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		fmt.Fprintln(rw, host)
		fmt.Fprintln(rw, r.URL.String())
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
