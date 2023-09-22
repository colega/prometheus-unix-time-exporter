package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	reg := prometheus.NewRegistry()
	reg.MustRegister(prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "unix_time",
		Help: "Current Unix time.",
	}, func() float64 { return float64(time.Now().Unix()) }))

	log.Println("listening on :8080")
	handler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
