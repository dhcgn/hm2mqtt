package server

import (
	"fmt"
	"net/http"
	"time"
)

//New create a http server
func New(mux *http.ServeMux, port int) *http.Server {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
	return srv
}
