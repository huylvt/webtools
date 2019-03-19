package main

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	"fmt"
	"strings"
)


// HTTP header check
func HTTPHeaderCheck(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	fmt.Println(url)

	resp, err := http.Head(url)
	
	if err != nil {
		log.Fatal(err)
	}

	m := make(map[string]string)
	m["Proto"] = resp.Proto
	m["Status"] = resp.Status
	
	for k, v := range resp.Header {
		m[k] = strings.Join(v, " ")
	}
	json.NewEncoder(w).Encode(m)
}

func main() {
	
	router := mux.NewRouter()
	router.HandleFunc("/http", HTTPHeaderCheck).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}