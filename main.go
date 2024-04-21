package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	log.Print("main.main")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("main.main HandleFunc")
		fmt.Fprintf(w, "Hello, %q!", html.EscapeString(r.URL.Path))
	})
	log.Print("main.main ListenAndServe: ", http.ListenAndServe(fmt.Sprintf(":%d", must2(strconv.Atoi(os.Getenv("PORT")))), nil))
}

func must2[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
