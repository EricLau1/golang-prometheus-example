package client

import (
	"flag"
	"fmt"
	"golang-prometheus-example/app/random"
	"net/http"
	"net/url"
	"time"
)

var (
	baseURL string
)

func init() {
	flag.StringVar(&baseURL, "base_url", "http://localhost:8080", "api base url")
	flag.Parse()
}

func Run() {
	statuses := []int{
		http.StatusOK,
		http.StatusCreated,
		http.StatusAccepted,
		http.StatusBadRequest,
		http.StatusNotFound,
		http.StatusUnauthorized,
		http.StatusTeapot,
		http.StatusInternalServerError,
	}
	for {
		delay := random.Intn(2000)
		for i := 0; i < len(statuses); i++ {
			status := statuses[random.Intn(len(statuses))]
			query := &url.Values{}
			query.Set("status", fmt.Sprint(status))
			go Request(query)
		}
		time.Sleep(time.Millisecond * time.Duration(delay))
	}
}
