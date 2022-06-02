package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service")

	// TODO connect to db

	// set up Config
	app := Config{}

	serve := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := serve.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
