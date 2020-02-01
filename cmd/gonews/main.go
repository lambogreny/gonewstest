package main

import (
	"net/http"

	"test/go/src/github.com/lambogreny/gonews/pkg/app"
)

const port = ":8080"

func main() {

	mux := http.NewServeMux()
	app.Mount(mux)
	http.ListenAndServe(port, mux)
	// Simple static webserver:

}
