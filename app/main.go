package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ecosia/go-prometheus-workshop/app/fetch"
)

const (
	port = "8000"
)

func handler(w http.ResponseWriter, r *http.Request) {
	statusCode, err := fetch.Fetch(fetch.NewRequest)

	if err == nil && statusCode == http.StatusOK {
		// Takes Tree Data from fetch package and marshal's it to the resp
		resp, err := json.Marshal(fetch.TreeData.Count)
		if err != nil {
			w.WriteHeader(500)
		}
		w.Write(resp)
	} else if statusCode != http.StatusOK {
		w.WriteHeader(statusCode)
	} else {
		w.WriteHeader(500)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	fmt.Printf("Service started at %v", port)
	http.ListenAndServe("0.0.0.0:"+port, mux)
}
