package main

import (
	"github.com/golang/glog"
	"flag"
	"os"
	"net/http"
	"time"
)

var (
	clientName = "client"
	serverURL = "http://localhost:8888/status"
	wait = 500
)

func setup() {
	flag.Set("logtostderr", "true")
	flag.Parse()
	su := os.Getenv("SERVER_URL")
	if su != "" {
		serverURL = su
		glog.Infof("SERVER_URL=%s", serverURL)
	}
}

func main() {
	setup()
	glog.Infof("Starting Client Name [%s]", clientName)
	for {
		request, _ := http.NewRequest("GET", serverURL, nil)
		client := &http.Client{}
		response, _ := client.Do(request)
		glog.Infof("[%s] Request to %s - Code [%d]", clientName, serverURL, response.StatusCode)
		requestSleep := time.Duration(wait) * time.Millisecond
		time.Sleep(requestSleep)
	}
}
