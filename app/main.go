package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"html/template"
	"net/http"

	"github.com/ecosia/go-prometheus-workshop/app/fetch"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	port = "8000"
)

var requestCounter = prometheus.NewCounter(prometheus.CounterOpts{Name: "requests_total"})

func handler(w http.ResponseWriter, r *http.Request) {
	requestCounter.Inc()

	t, err := template.ParseFiles("./templates/withResponse.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	treeData, statusCode, err := fetch.Fetch(fetch.NewRequest)
	if statusCode != http.StatusOK {
		w.WriteHeader(statusCode)
		return
	}

	err = t.Execute(w, treeData.Count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	// metrics
	prometheus.MustRegister(requestCounter)
	mux.Handle("/metrics", promhttp.Handler())

	fmt.Printf("Service started at %v", port)
	http.ListenAndServe("0.0.0.0:"+port, mux)
}
