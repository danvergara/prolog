package main

import (
	"log"

	"github.com/danvergara/prolog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8000")
	log.Fatal(srv.ListenAndServe())
}
