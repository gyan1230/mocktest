package main

import (
	"log"
	"mocktest/handler"
	"net/http"
	"time"
)

func main() {
	httpClient := &http.Client{Timeout: time.Second * 60}
	srvHandle := &handler.Mock{HttpClient: httpClient,
		CurrSrvURL: "https://api.coindesk.com/v1/bpi/currentprice.json"}
	log.Println("Starting http service...")
	if err := http.ListenAndServe(":8001", srvHandle); err != nil {
		log.Fatal("Error :", err)
	}
}
