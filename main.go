package main

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/charliekim2/progress/views/layouts"
)

func handler(w http.ResponseWriter, r *http.Request) {
	helloLayout := layouts.Hello(r.URL.Path[1:])
	render(helloLayout, w)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func render(c templ.Component, w io.Writer) {
	c.Render(context.Background(), w)
}
