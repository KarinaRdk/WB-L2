package main

import (
	"net/http"

)

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)
	mux.HandleFunc("/v1/new", app.createEvent)
	mux.HandleFunc("/v1/new:id", app.getEvent)

	return mux
	
}