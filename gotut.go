package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("d:/logs/gotuts.log",
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
	defer f.Close()
	url := "http://ftp.funet.fi"
	for {
		resp, err := http.Get(url)
		if err != nil {
			log.Print(url, " failed: ", err)
		}
		resp.Body.Close()
		time.Sleep(1 * time.Second)
	}
}
