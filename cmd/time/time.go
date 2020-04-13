package main

import (
	"flag"
	"log"
	"os"
	"time"

	"encoding/json"
	"net/http"
)

type timeResponse struct {
	Hostname string
	Time     string
}

func TimeHandler(hostname string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("[%s] - [%s] - %s", req.RemoteAddr, req.Proto, req.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(timeResponse{
			Hostname: hostname,
			Time:     time.Now().Format(time.Kitchen),
		})
	})
}

func main() {
	addr := flag.String("addr", ":9000", "listen addr")
	flag.Parse()

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/time", TimeHandler(hostname))
	log.Fatal(http.ListenAndServe(*addr, nil))
}
