package main

import (
	"html/template"
	"net/http"
)

type response struct {
	count int
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := response{count: 1}
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
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	http.ListenAndServe("0.0.0.0:8000", mux)
}
