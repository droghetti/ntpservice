package main

import (
    "fmt"
	"time"
    "net/http"
)


func _unixtime(w http.ResponseWriter, r *http.Request) {
    t := time.Now()
	fmt.Fprintf(w, "%du", t.UTC().Unix())
}

func _utctime(w http.ResponseWriter, r *http.Request) {
    t := time.Now()
	fmt.Fprintf(w,"%s", t.UTC())
}

func main() {
    http.HandleFunc("/", _unixtime)
    http.HandleFunc("/utctime", _utctime)
    http.ListenAndServe(":8081", nil)
}
