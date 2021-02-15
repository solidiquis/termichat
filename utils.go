package main

import (
	"net/http"
	"strings"
)

func ipAddress(r *http.Request) string {
	header := r.Header
	realIP := header.Get("X-Real-Ip")
	forwardedFor := header.Get("X-Forwarded-For")

	if realIP == "" && forwardedFor == "" {
		idx := strings.LastIndex(r.RemoteAddr, ":")

		if idx == -1 {
			return r.RemoteAddr
		}

		return r.RemoteAddr[:idx]
	}

	if forwardedFor != "" {
		parts := strings.Split(forwardedFor, ",")
		for i, p := range parts {
			parts[i] = strings.TrimSpace(p)
		}

		return parts[0]
	}

	return realIP
}
