package main

import (
	"flag"
	"fmt"
	"log"
	"net"
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
		addr, err := ip()
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
		fmt.Fprintf(rw, "%s\n", addr)
		fmt.Fprintf(rw, "\tURL=%s", r.URL.String())
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

func ip() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "N/A", nil
}
