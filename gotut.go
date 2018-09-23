package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	f, err := os.OpenFile("d:/logs/gotuts.log",
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(&lumberjack.Logger{
		Filename:   "d:/logs/gotuts.log",
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     28,    //days
		Compress:   false, // disabled by default
	})
	timeout := time.Duration(10 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}
	defer f.Close()
	urls := []string{"http://ftp.funet.fi",
		"http://157.24.8.10"}
	for {
		for _, url := range urls {
			resp, err := client.Get(url)
			if err != nil {
				log.Print(url, " failed: ", err)
			} else {
				resp.Body.Close()
			}
		}
		time.Sleep(1 * time.Second)
	}
}
