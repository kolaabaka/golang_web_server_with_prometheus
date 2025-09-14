package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	promCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "myCoiunter",
			Help: "Help",
		},
		[]string{"label_name"},
	)
)

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware handler")
		next.ServeHTTP(w, r)
	})
}

func prometheusCounterStat(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		promCounter.WithLabelValues("").Inc()
		next.ServeHTTP(w, r)
	})
}

func resp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi, my friend"))
}

func main() {

	prometheus.MustRegister(promCounter)

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/", loggerMiddleware(prometheusCounterStat(http.HandlerFunc(resp))))
	http.ListenAndServe(":8080", nil)
}
