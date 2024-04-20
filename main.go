package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q!", html.EscapeString(r.URL.Path))
	})
	http.ListenAndServe(fmt.Sprintf(":%d", must2(strconv.Atoi(os.Getenv("PORT")))), nil)
}

func must2[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
