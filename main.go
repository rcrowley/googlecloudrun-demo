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
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}
	log.Print("main.main ListenAndServe: ", http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
