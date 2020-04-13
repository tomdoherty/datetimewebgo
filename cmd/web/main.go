package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"encoding/json"
	"net/http"
)

type webResponse struct {
	Hostname   string `json:"hostname"`
	Time       string `json:"time"`
	TimeSource string `json:"timesource"`
	Date       string `json:"date"`
	DateSource string `json:"datesource"`
}

func main() {
	addr := flag.String("addr", ":7000", "listen addr")

	dateaddr := flag.String("dateaddr", "date:8000", "date service addr")
	timeaddr := flag.String("timeaddr", "time:9000", "time service addr")

	flag.Parse()

	dateurl := "http://" + *dateaddr + "/date"
	timeurl := "http://" + *timeaddr + "/time"

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", getDateAndTime(hostname, dateurl, timeurl))
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func getDateAndTime(hostname, dateurl, timeurl string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("[%s] - [%s] - %s", req.RemoteAddr, req.Proto, req.URL.Path)
		w.Header().Set("Content-Type", "application/json")

		dateHostname, dateResponse, err := getDate(dateurl)
		if err != nil {
			log.Printf("failed to getDate(): %w", err)
			return
		}

		timeHostname, timeResponse, err := getTime(timeurl)
		if err != nil {
			log.Printf("failed to getTime(): %w", err)
			return
		}

		json.NewEncoder(w).Encode(webResponse{
			Hostname:   hostname,
			Date:       dateResponse,
			DateSource: dateHostname,
			Time:       timeResponse,
			TimeSource: timeHostname,
		})
	})
}

func getDateOrTime(dateurl string) (response webResponse, err error) {
	resp, err := http.Get(dateurl)
	if err != nil {
		return response, fmt.Errorf("failed to fetch: %w", err)
	}
	defer resp.Body.Close()

	var result webResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return response, fmt.Errorf("failed to decode: %w", err)
	}
	return result, nil
}

func getDate(dateurl string) (hostname, date string, err error) {
	result, err := getDateOrTime(dateurl)
	if err != nil {
		return "", "", fmt.Errorf("failed to getDate: %w", err)
	}
	return result.Hostname, result.Date, nil
}

func getTime(timeurl string) (hostname, time string, err error) {
	result, err := getDateOrTime(timeurl)
	if err != nil {
		return "", "", fmt.Errorf("failed to getTime: %w", err)
	}
	return result.Hostname, result.Time, nil
}
