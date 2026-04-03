package main

import (
	"log"
	"net/http"
	"os"

	"github.com/pmenglund/gui/examples/showcase/app"
)

func main() {
	root, err := os.Getwd()
	if err != nil {
		log.Fatalf("determine working directory: %v", err)
	}

	addr := os.Getenv("PORT")
	if addr == "" {
		addr = "8080"
	}

	log.Printf("showcase listening on http://localhost:%s", addr)
	if err := http.ListenAndServe(":"+addr, app.NewMux(root)); err != nil {
		log.Fatal(err)
	}
}
