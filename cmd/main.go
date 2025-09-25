package main

import (
	"golang_web_server_with_prometheus/internal/middleware"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func resp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi, my friend"))
}

func main() {
	middleware.InitPromeheusStat()
	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/", middleware.LoggerMiddleware(middleware.PrometheusCounterStat(http.HandlerFunc(resp))))
	http.ListenAndServe(":8080", nil)
}
