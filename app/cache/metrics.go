package cache

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"log"
)

const defaultGatewayURL = "http://localhost:9091"

var (
	memoryGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cache_memory_usage",
		Help: "Quantidade memoria usada",
	})
)

func init() {
	prometheus.MustRegister(memoryGauge)
}

func Push() {
	pusher := push.New(defaultGatewayURL, "cache")
	err := pusher.Collector(memoryGauge).Push()
	if err != nil {
		log.Println("error on push metrics:", err.Error())
	} else {
		log.Println("push metrics successfully!")
	}
}
