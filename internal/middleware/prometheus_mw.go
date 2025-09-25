package middleware

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
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

func InitPromeheusStat() {
	prometheus.MustRegister(promCounter)
}

func PrometheusCounterStat(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		promCounter.WithLabelValues("").Inc()
		next.ServeHTTP(w, r)
	})
}
