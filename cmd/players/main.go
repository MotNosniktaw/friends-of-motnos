package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var counter = 0

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.HandleFunc("/counter", func(w http.ResponseWriter, r *http.Request) {
		counter++
		w.Write([]byte(strconv.Itoa(counter)))
	})

	fmt.Println("running players service")
	http.ListenAndServe(":50052", r)
}
