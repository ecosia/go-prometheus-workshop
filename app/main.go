package main

import (
	"encoding/json"
	"net/http"
)

type response struct {
	count int
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := response{count: 1}
	resp, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(500)
	}
	w.Write(resp)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	http.ListenAndServe("0.0.0.0:8000", mux)
}
