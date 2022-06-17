package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct {
}

const webPort = "80"

func main() {
	app := Config{}

	log.Println("Starting mail service on port", webPort)

	serve := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := serve.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
