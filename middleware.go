package main

import (
	"net/http"
)

func logRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		uri := r.URL.String()
		referer := r.Header.Get("Referer")
		userAgent := r.Header.Get("User-Agent")
		ipAddress := ipAddress(r)

		infoLog.Printf("%s %s %s %s %s\n",
			ipAddress,
			method,
			uri,
			referer,
			userAgent,
		)

		next(w, r)
	}
}
