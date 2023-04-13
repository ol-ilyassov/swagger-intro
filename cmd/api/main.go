package main

import (
	"log"
	"os"
)

const version = "1.0.0"

type config struct {
	env string
}

type application struct {
	config config
	logger *log.Logger
}

//	@title			Swagger Intro API
//	@version		1.0
//	@description	Swagger API for Golang Project Blueprint.

//	@contact.name	API Support
//	@contact.email	helpdesk@gmail.com

//	@license.name	MIT
//	@license.url	https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE

//	@BasePath	/api/v1
func main() {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		logger: logger,
	}

	// Start the HTTP server.
	logger.Printf("starting %s server on %s", "development", ":4000")
	err := app.serve()
	if err != nil {
		logger.Fatal(err)
	}
}
