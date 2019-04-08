package main

import (
	"html/template"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type response struct {
	count int
}

var (
	dummyCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "requests",
		},
	)
)

func handler(w http.ResponseWriter, r *http.Request) {
	data := response{count: 1}
	dummyCounter.Inc()
	t, err := template.ParseFiles("./templates/withResponse.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = t.Execute(w, data.count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func main() {
	prometheus.MustRegister(dummyCounter)
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe("0.0.0.0:8000", mux)
}
