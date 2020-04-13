package main

import (
	"flag"
	"log"
	"os"
	"time"

	"encoding/json"
	"net/http"
)

type dateResponse struct {
	Hostname string
	Date     string
}

func DateHandler(hostname string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("[%s] - [%s] - %s", req.RemoteAddr, req.Proto, req.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dateResponse{
			Hostname: hostname,
			Date:     time.Now().Format("2006/01/02"),
		})
	})
}

func main() {
	addr := flag.String("addr", ":8000", "listen addr")
	flag.Parse()

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/date", DateHandler(hostname))
	log.Fatal(http.ListenAndServe(*addr, nil))
}
