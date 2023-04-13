package main

import (
	"errors"
	"net/http"
	"time"
)

// Launches application as server.
func (app *application) serve() error {

	srv := &http.Server{
		Addr:    ":4000",
		Handler: app.routes(),
		// ErrorLog:     log.New(app.logger, "", 0), // ? extra functionality
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	app.logger.Print("stopped server", map[string]string{
		"addr": srv.Addr,
	})

	return nil
}
