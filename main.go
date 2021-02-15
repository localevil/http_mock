package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Req struct {
	Status   int
	Response string
}

type ReqAdd struct {
	Url     string
	Request Req
}

type requests map[string]Req

func main() {
	m := make(requests)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		req, ok := m[r.URL.Path]
		log.Println(r.URL.Path)
		log.Println(m)
		if !ok {
			http.Error(w, "Not Found", http.StatusNotFound)
		} else if req.Status != 200 {
			http.Error(w, "", req.Status)
		} else {
			fmt.Fprintf(w, req.Response)
		}
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		var req ReqAdd
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Fprintf(w, err.Error(), http.StatusBadRequest)
		}
		log.Println(req)
		m[req.Url] = req.Request
		fmt.Fprintf(w, "OK")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
