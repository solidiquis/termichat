package main

import (
	"fmt"
	"net/http"
)

var (
	server *http.Server
)

func init() {
	server = &http.Server{
		Addr:           fmt.Sprintf(":%d", DEFAULT_PORT),
		Handler:        router,
		ErrorLog:       errLog,
		MaxHeaderBytes: 1 << 20,
	}
}

func termichat() {
	infoLog.Printf("Listening on port %d\n", DEFAULT_PORT)
	errLog.Fatal(server.ListenAndServe())
}
