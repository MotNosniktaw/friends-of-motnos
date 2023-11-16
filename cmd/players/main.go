package main

import (
	"fmt"
	"net/http"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	fmt.Println("running players service")
	http.ListenAndServe(":50052", r)
}
