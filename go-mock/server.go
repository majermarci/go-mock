package main

import (
	"encoding/json"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	port    = flag.String("p", "8080", "Port where the server will listen (default: 8080)")
	version = "0.5.6"
)

// Serve is called with loaded config by main func
func serve(c config) {
	// Reserved health probe endpoint without logging
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{
			"status":  "running",
			"version": version,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			slog.Error("error writing health response", "details", err)
		}
	})

	http.HandleFunc("/paths", func(w http.ResponseWriter, r *http.Request) {

		adminPassword := "admin"
		if envPass := os.Getenv("MOCK_ADMIN_PASS"); envPass != "" {
			adminPassword = envPass
		}

		username, password, ok := r.BasicAuth()
		if !ok || username != "admin" || password != adminPassword {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			slog.Warn("wrong credentials used on /paths endpoint.")
			return
		}

		pathsResponse := map[string][]string{}

		for path, methods := range c.Endpoints {
			for method := range methods {
				pathsResponse[method] = append(pathsResponse[method], path)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(pathsResponse)
		if err != nil {
			slog.Error("error writing paths response", "details", err)
		}

		slog.Info("paths endpoint requested")
	})

	// Default handler for all other routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		method := r.Method

		if endpoint, ok := c.Endpoints[path][method]; ok {
			w.Header().Set("System", "mock-server")

			for key, value := range endpoint.Headers {
				w.Header().Set(key, value)
			}

			slog.Info("request mocked",
				"path", path,
				"method", method,
				"headers_received", r.Header,
				"headers_sent", w.Header(),
			)

			body := strings.TrimSpace(endpoint.Body)
			if isJSON(body) {
				formattedBody, err := formatJSON(body)
				if err != nil {
					slog.Error("invalid JSON in response body", "details", err)
					http.Error(w, "invalid JSON in server config", http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				body = formattedBody
			}

			w.Header().Set("Content-Length", strconv.Itoa(len(body)))
			w.WriteHeader(endpoint.Status)

			if body != "" {
				_, err := w.Write([]byte(body))
				if err != nil {
					slog.Error("error writing response", "details", err)
				}
			}
		} else {
			slog.Error("not found", "method", method, "path", path)
			http.NotFound(w, r)
		}
	})

	_, err := strconv.Atoi(*port)
	if err != nil {
		slog.Error("invalid port:"+*port, "details", err)
		os.Exit(1)
	}

	slog.Info("Starting server on port " + *port)
	err = http.ListenAndServe(":"+*port, nil)
	if err != nil {
		slog.Error("error with server", "details", err)
		os.Exit(1)
	}
}
