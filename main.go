// Notify Service - © 2022 SouthWinds Tech Ltd - www.southwinds.io
// Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
// Contributors to this project, hereby assign copyright in this code to the project,
// to be licensed under the same terms as the rest of the code.
package main

import (
	"github.com/gorilla/mux"
	h "southwinds.dev/http"
	notify "southwinds.dev/notify/client"
	"southwinds.dev/notify/core"
)

func main() {
	// creates a generic http server
	s := h.New("notify", notify.Version)
	// add handlers
	s.Http = func(router *mux.Router) {
		if core.IsLoggingEnabled() {
			router.Use(s.LoggingMiddleware)
		}
		router.Use(s.AuthenticationMiddleware)
		router.HandleFunc("/notify", notifyHandler).Methods("POST")
	}
	s.Serve()
}
